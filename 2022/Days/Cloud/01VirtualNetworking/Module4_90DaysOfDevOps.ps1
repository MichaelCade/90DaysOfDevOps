$rgName = '90DaysOfDevOps'

New-AzResourceGroupDeployment `
-ResourceGroupName $rgName `
-TemplateFile C:\Users\micha\demo\90DaysOfDevOps\Days\Cloud\01VirtualNetworking\Mod04_90DaysOfDevOps-vms-loop-template.json `
-TemplateParameterFile C:\Users\micha\demo\90DaysOfDevOps\Days\Cloud\01VirtualNetworking\Mod04_90DaysOfDevOps-vms-loop-parameters.json