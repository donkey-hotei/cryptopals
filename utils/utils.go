package utils

import (
    b64 "encoding/base64"
    "encoding/hex"
    "math"
)

func check(err error) { 
    if err != nil {
        panic(err)
    }
}

// https://en.wikipedia.org/wiki/Letter_frequency
var englishFreqs = []float64 {
    //   a       b       c       d      e       f       g
    0.0817, 0.0149, 0.0278, 0.0425, 0.127, 0.0223, 0.0202,
    //   h       i       j       k      l       m       n
    0.0609, 0.0697, 0.0015, 0.0077, 0.040, 0.0241, 0.0675,
    //   o       p       q       r      s       t       u
    0.0751, 0.0193, 0.0095, 0.0599, 0.063, 0.0906, 0.0276,
    //   v       w       x       y      z  SPACE
    0.0098, 0.0236, 0.0015, 0.0197, 0.007, 0.191,
}

/*
 * Computes u[0] * v[0] + ... + u[n] + v[n]
 * where n is the length of the smallest vector.
 */
func DotProduct(u, v []float64) float64 {
    dot_product := 0.0

    for i := range u {
        dot_product += u[i] * v[i]
    }

    return dot_product
}

/*
 * Computes the magnitude of a vector u.
 */
func LenVec(u []float64) float64 {
    return math.Sqrt(DotProduct(u, u))
}

/*
 * Computes the Hamming distance of two vectors a and b.
 */
func HammingDistance(a, b []byte) int {
    distance := 0

    for i, _ := range a {
        if a[i] != b[i] {
            distance += 1
            println(distance)
        }
    }

    return distance
}

/*
 * Computes the cosine similarity of two vectors a and b.
 */
func Cosine(a, b []float64) float64 {
    return DotProduct(a, b) / (LenVec(a) * LenVec(b))
}

/*
 * Scores a text based on it's cosine similarity with a vector of ideal
 * English letter frequencies.
 */
func ScoreText(s []byte) float64 {
    counts := make([]int, 27)

    for _, ch := range s {
        if 'A' <= ch && ch <= 'Z' {
            counts[ch - 65]++
        }
        if 'a' <= ch && ch <= 'z' {
            counts[ch - 97]++
        }
        if ch == ' ' {
            counts[26]++
        }
    }

    score := 0.0
    freqs := make([]float64, 27)
    amt   := float64(len(s))

    for i, c := range counts {
        freqs[i] = float64(c) / amt
        score   += freqs[i]
    }

    return Cosine(freqs, englishFreqs)
}

func HexToBase64(data string) []byte {
    str, err := hex.DecodeString(data)

    check(err)

    b64_string := make([]byte, b64.StdEncoding.EncodedLen(len(str)))
    b64.StdEncoding.Encode(b64_string, str)

    return b64_string
}

