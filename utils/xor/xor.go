package xor

import (
    "github.com/donkey-hotei/cryptopals/utils"
    "encoding/hex"
)


/*
 * Bruteforce key for a single byte xor cipher by finding the byte from
 * 0x0 to 0xff that has the highest score.
 */
func CrackSingleKeyXOR(ct []byte) int {
    result     := make([]byte, len(ct))
    best_score := 0.0
    likely_key := 0

    for k := 1; k < 255; k++ {
        result = ByteXOR(byte(k), ct)
        score := utils.ScoreText(result)

        if score >= best_score {
            best_score = score
            likely_key = k
        }
    }

    return likely_key
}

/*
 *
 */
func CrackRepeatingKeyXOR(ct []byte) []byte {
    return ct
}


/*
 * Finds and decrypts the most likely candidate to be encrypted
 * with single-character XOR.
 */
func DetectSingleKeyXOR(cts []string) (float64, string) {
    var best_result []byte

    best_score  := 0.0

    for _, line := range cts {
        ct, _ := hex.DecodeString(line)
        k     := CrackSingleKeyXOR(ct)
        pt    := ByteXOR(byte(k), ct)
        score := utils.ScoreText(pt)

        if score >= best_score {
            best_score  = score
            best_result = pt
        }
    }

    return best_score, string(best_result)
}

/*
 * Repeatedly XORs key against the plaintext pt.
 */
func RepeatedKeyXOR(key, pt []byte) []byte {
    result  := make([]byte, len(pt))
    key_len := len(key)

    for i, _ := range pt {
        result[i] = pt[i] ^ key[i % key_len]
    }

    return result
}

/*
 * XOR byte vectors a and b.
 */
func FixedXOR(a, b []byte) []byte {
    n := len(a)

    if len(b) < n {
        n = len(b)
    }

    dst := make([]byte, n)

    for i := 0; i < n; i++ {
        dst[i] = a[i] ^ b[i]
    }

    return dst
}

/*
 * XOR all bytes in b with byte k
 */
func ByteXOR(k byte, b []byte) []byte {
    result := make([]byte, len(b))

    for i := range b {
        result[i] = b[i] ^ k
    }

    return result
}


