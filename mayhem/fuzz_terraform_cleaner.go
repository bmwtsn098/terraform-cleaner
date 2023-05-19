package fuzz_terraform_cleaner

import (
    fuzz "github.com/AdaLogics/go-fuzz-headers"

    "github.com/sylwit/terraform-cleaner/terraform"
)

func mayhemit(data []byte) int {

    if len(data) > 2 {
        num := int(data[0])
        data = data[1:]
        fuzzConsumer := fuzz.NewConsumer(data)
        
        switch num {
            
            case 0:
                testPath, _ := fuzzConsumer.GetString()
                
                terraform.NewModuleUsage(testPath)
                return 0

            case 1:
                testPath, _ := fuzzConsumer.GetString()
                
                terraform.ListTfModules(testPath)
                return 0

            case 2:
                testPath, _ := fuzzConsumer.GetString()
                
                terraform.LoadTfModule(testPath)
                return 0

            case 3:
                var testModule terraform.ModuleUsage
                fuzzConsumer.GenerateStruct(&testModule)
                testBool, _ := fuzzConsumer.GetBool()

                testModule.DisplayLocals(testBool)
                return 0

            case 4:
                var testModule terraform.ModuleUsage
                fuzzConsumer.GenerateStruct(&testModule)
                testBool, _ := fuzzConsumer.GetBool()

                testModule.DisplayVariables(testBool)
                return 0

            case 5:
                var testModule terraform.ModuleUsage
                fuzzConsumer.GenerateStruct(&testModule)
                testBool, _ := fuzzConsumer.GetBool()
                temp, _ := fuzzConsumer.GetString()
                testType := terraform.DisplayType(temp)

                testModule.Display(testType, testBool)
                return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}