package main

func (bc *BC) scandirs(message string) error {
	return bc.Operation.ScanDirs(bc, message)
}
