package main

func (bc *BC) create(message string) error {
	return bc.Operation.Create(bc, message)
}
