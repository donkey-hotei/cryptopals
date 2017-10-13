package main

import (
    xor "github.com/donkey-hotei/cryptopals/utils/xor"
    "github.com/donkey-hotei/cryptopals/utils"
    "encoding/hex"
    "strings"
    "fmt"
//    "os"
//    "io"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func one_one() {
    // 1.1
    data   := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    result := utils.HexToBase64(data)
    fmt.Printf("%q\n", result)
}

func one_two() {
    // 1.2
    s1, _  := hex.DecodeString("1c0111001f010100061a024b53535009181c")
    s2, _  := hex.DecodeString("686974207468652062756c6c277320657965")
    result := xor.FixedXOR(s1, s2)
    fmt.Printf("%q\n", result)
}

func one_three() {
    // 1.3
    ct, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
    k     := xor.CrackSingleKeyXOR(ct)
    pt    := xor.ByteXOR(byte(k), ct)
    fmt.Printf("%q\n", string(pt))
}

func one_four() {
    // 1.4
    data, _ := ioutil.ReadFile("./data/4.txt")
    file_contents := strings.Split(string(data), "\n")
    score, plaintext := xor.DetectSingleKeyXOR(file_contents)
    fmt.Printf("%f: %q\n", score, plaintext)
}

func one_five() {
    // 1.5
    pt  := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal\n")
    key := []byte("ICE")
    ct  := xor.RepeatedKeyXOR(key, pt)
    fmt.Println(hex.EncodeToString(ct))
}

func main() {
    one_one()
    one_two()
    one_three()
    one_four()
    one_five()

    string_one := []byte("this is a test\n")
    string_two := []byte("wokka wokka!!!\n") // should equal 37
    distance   := utils.HammingDistance(string_one, string_two)
    print(distance)
}