package utils

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/manifoldco/promptui"
)

func GetAddress() (address common.Address, err error) {
	return GetAddressWithPrompt("Enter Address")
}

func GetAddressWithPrompt(promptLabel string) (address common.Address, err error) {
	validate := func(input string) error {
		isValid := common.IsHexAddress(input)
		if !isValid {
			return errors.New("invalid EVM Address")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    promptLabel,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return common.Address{}, err
	}

	return common.HexToAddress(result), nil
}

func GetAmountWithPrompt(promptLabel string) (*big.Float, error) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid amount")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    promptLabel,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	amount, _, err := big.ParseFloat(result, 10, 0, big.ToZero)
	if err != nil {
		return nil, err
	}

	amount.SetPrec(64)
	return amount, err
}

func GetIntWithPrompt(promptLabel string) (int, error) {
	validate := func(input string) error {
		_, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("invalid amount")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    promptLabel,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0, err
	}

	amount, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return amount, err
}

func GetAmount() (weiAmount *big.Float, err error) {
	return GetAmountWithPrompt("Enter Amount")
}

func GetStringWithPrompt(promptLabel string) (string, error) {
	prompt := promptui.Prompt{
		Label: promptLabel,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, err
}

func GetHexWithPrompt(promptLabel string) (str string, err error) {
	//validate := func(s string) error {
	//	pattern := `^(0x)?[0-9a-fA-F]*$`
	//	matched, _ := regexp.MatchString(pattern, s)
	//	if matched {
	//		return nil
	//	}
	//	return errors.New("not a valid hex string")
	//}

	//prompt := promptui.Prompt{
	//	Label: promptLabel,
	//	Validate: validate,
	//}

	//str, err = prompt.Run()
	//if err != nil {
	//	fmt.Printf("Prompt failed %v\n", err)
	//	return
	//}

	fmt.Printf("%s: ", promptLabel)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if err = scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	return input, err
}

func SelectSingleOption(promptTitle string, options []string) (selection string, err error) {
	var answer core.OptionAnswer
	prompt := &survey.Select{
		Message: promptTitle,
		Options: options,
	}
	err = survey.AskOne(prompt, &answer)
	if err != nil {
		fmt.Printf("failed selection: %s\n", err)
		return
	}

	return answer.Value, nil
}
