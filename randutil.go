/*
                                    Goutil
             A Motley Assortment of Utilities for the Go Language


@author: Jason McVetta <jason.mcvetta@gmail.com>
@copyright: (c) 2012 Jason McVetta
@license: GPL v3 - http://www.gnu.org/copyleft/gpl.html

********************************************************************************
Goutil is free software: you can redistribute it and/or modify it under the
terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

Goutil is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
Goutil.  If not, see <http://www.gnu.org/licenses/>.
********************************************************************************

*/

package goutil

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

const (
	// Set of characters to use for generating random strings
	Alphanumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 abcdefghijklmnopqrstuvwxyz"
	Ascii        = Alphanumeric + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

// RandString returns a random string no more than at least min and no more 
// than max characters long.
func RandString(min, max int, charset string) string {
	//
	// First determine the length of string to be generated
	//
	var err error  // Holds errors
	var b *big.Int // Holds random bigints
	var r int      // Holds random integers
	var strlen int // Length of random string to generate
	switch {
	case max == min:
		strlen = max
	case max > min:
		// Choose a random string lenth between min and max
		maxRand := max - min
		b, err = rand.Int(rand.Reader, big.NewInt(int64(maxRand)))
		if err != nil {
			panic(err) // WTF?
		}
		r = int(b.Int64())
		strlen = min + r
	case min < max:
		msg := "Min (%s) cannot be greater than max (%s)!"
		msg = fmt.Sprintf(msg, min, max)
		log.Panic(msg)
	}
	//
	// Generate a random string that is strlen characters long
	//
	charlen := len(charset)
	randstr := ""
	for i := 0; i < strlen; i++ {
		b, err = rand.Int(rand.Reader, big.NewInt(int64(charlen-1)))
		if err != nil {
			panic(err) // WTF?
		}
		r = int(b.Int64())
		c := string(charset[r])
		randstr += c
	}
	return randstr
}

func RandAlphanumeric(min, max int) string {
	return RandString(min, max, Alphanumeric)
}