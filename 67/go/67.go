import "math/big"

func addBinary(a string, b string) string {
    numA := new(big.Int)
    numB := new(big.Int)
    numR := new(big.Int)

    numA.SetString(a, 2)
    numB.SetString(b, 2)

    numR.Add(numA, numB)

    return numR.Text(2)
}