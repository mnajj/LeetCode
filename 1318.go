func minFlips(a int, b int, c int) int {
    flips := 0
    for a > 0 || b > 0 || c > 0 {
        abit, bbit, cbit := a & 1, b & 1, c & 1
        if abit | bbit != cbit {
            flips += abit + bbit + cbit
        }
        a, b, c = a >> 1, b >> 1, c >> 1
    }
    return flips
}
