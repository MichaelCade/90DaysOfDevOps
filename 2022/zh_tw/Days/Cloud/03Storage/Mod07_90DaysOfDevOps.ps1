$rgName = '90DaysOfDevOps'

New-AzResourceGroupDeployment `
   -ResourceGroupName $rgName `
   -TemplateFile C:\Users\micha\demo\90DaysOfDevOps\Days\Cloud\03Storage\Mod07_90DaysOfDevOps-vm-template.json `
   -TemplateParameterFile C:\Users\micha\demo\90DaysOfDevOps\Days\Cloud\03Storage\Mod07_90DaysOfDevOps-vm-parameters.json `
   -AsJob
