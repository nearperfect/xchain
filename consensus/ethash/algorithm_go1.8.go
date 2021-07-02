// Copyright 2017 The MOAC-core Authors
// This file is part of the MOAC-core library.
//
// The MOAC-core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The MOAC-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the MOAC-core library. If not, see <http://www.gnu.org/licenses/>.

// +build go1.8

package ethash

import "math/big"

// cacheSize calculates and returns the size of the ethash verification cache that
// belongs to a certain block number. The cache size grows linearly, however, we
// always take the highest prime below the linearly growing threshold in order to
// reduce the risk of accidental regularities leading to cyclic behavior.
func cacheSize(block uint64) uint64 {
	// If we have a pre-generated value, use that
	epoch := int(block / epochLength)
	if epoch < maxEpoch {
		return cacheSizes[epoch]
	}

	return calcCacheSize(epoch)
}

// calcCacheSize calculates the cache size for epoch. The cache size grows linearly,
// however, we always take the highest prime below the linearly growing threshold in order
// to reduce the risk of accidental regularities leading to cyclic behavior.
func calcCacheSize(epoch int) uint64 {
	size := cacheInitBytes + cacheGrowthBytes*uint64(epoch) - hashBytes
	for !new(big.Int).SetUint64(size / hashBytes).ProbablyPrime(1) { // Always accurate for n < 2^64
		size -= 2 * hashBytes
	}
	return size
}

// datasetSize calculates and returns the size of the ethash mining dataset that
// belongs to a certain block number. The dataset size grows linearly, however, we
// always take the highest prime below the linearly growing threshold in order to
// reduce the risk of accidental regularities leading to cyclic behavior.
func datasetSize(block uint64) uint64 {
	// If we have a pre-generated value, use that
	epoch := int(block / epochLength)
	if epoch < maxEpoch {
		return datasetSizes[epoch]
	}

	return calcDatasetSize(epoch)
}

// calcDatasetSize calculates the dataset size for epoch. The dataset size grows linearly,
// however, we always take the highest prime below the linearly growing threshold in order
// to reduce the risk of accidental regularities leading to cyclic behavior.
func calcDatasetSize(epoch int) uint64 {
	size := datasetInitBytes + datasetGrowthBytes*uint64(epoch) - mixBytes
	for !new(big.Int).SetUint64(size / mixBytes).ProbablyPrime(1) { // Always accurate for n < 2^64
		size -= 2 * mixBytes
	}
	return size
}
