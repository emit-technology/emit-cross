// Copyright 2020 EMIT Foundation.
// This file is part of E.M.I.T. .

// E.M.I.T. is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// E.M.I.T. is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with E.M.I.T. . If not, see <http://www.gnu.org/licenses/>.

package queue

import (
	"errors"
)

var (

	// ErrEmpty is returned when the stack or queue is empty.
	ErrEmpty = errors.New("queue is empty")

	// ErrOutOfBounds is returned when the ID used to lookup an item
	// is outside of the range of the stack or queue.
	ErrOutOfBounds = errors.New("ID used is outside range of queue")

	// ErrDBClosed is returned when the Close function has already
	// been called, causing the stack or queue to close, as well as
	// its underlying database.
	ErrDBClosed = errors.New("Database is closed")
)
