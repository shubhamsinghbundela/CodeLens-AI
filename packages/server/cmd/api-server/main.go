package main

import "my-codelens-app/internal/common"

func main() {
	ctx, cancel := common.GlobalContext()
	defer cancel()
}
