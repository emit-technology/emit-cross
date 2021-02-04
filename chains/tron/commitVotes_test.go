// Copyright 2020 The emit-cross Authors
// This file is part of the emit-cross library.
//
// The emit-cross library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The emit-cross library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the emit-cross library. If not, see <http://www.gnu.org/licenses/>.

package tron

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeLoop(t *testing.T) {
	fmt.Println("watchInterval", uint64(WatchDuration))

	fmt.Println("begin", time.Now().Unix())

	time.Sleep(time.Second * 2)

	fmt.Println("end", time.Now().Unix())

}
