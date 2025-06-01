package valobj

import "testing"

func TestPasswordNew(t *testing.T) {
	valids := []string{
		"Valid1@Password",
		"123456aZ@",
		"Aa1@345678",
		"longPassword2@",
	}

	invalids := []string{
		"Short@1",
		"validButTooLongPassword123456789!@#$%",
		"lowercaseonly",
		"UPPERCASEONLY",
		"ｚｅｎｋａｋｕｏｎｌｙ",
		"123456",
		"onlyLetters@",
		"contain space",
		"contain	tab",
		"notContainSymbol123",
		"notContainNumber!@",
		"NOTCONTAINLOWER!@12",
		"notcontainupper!@12",
		"123456!@#$",
	}

	for _, password := range valids {
		_, err := NewPassword(password)

		if err != nil {
			t.Errorf("expected valid password: %s", password)
		}
	}

	for _, password := range invalids {
		_, err := NewPassword(password)

		if err == nil {
			t.Errorf("expected invalid password: %s", password)
		}
	}
}

func TestPasswordHashCheck(t *testing.T) {
	plain := "Password@123"
	p, _ := NewPassword(plain)

	hashed, err := p.Hash()
	if err != nil {
		t.Errorf("password hash was failed, %v", err)
	}

	if !p.Check(hashed) {
		t.Errorf("hash check wa failed")
	}
}
