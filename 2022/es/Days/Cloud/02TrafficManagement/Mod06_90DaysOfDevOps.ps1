$rgName = '90DaysOfDevOps'

New-AzResourceGroupDeployment `
   -ResourceGroupName $rgName `
   -TemplateFile C:\Users\micha\demo\90DaysOfDevOps\Days\Cloud\02TrafficManagement\Mod06_90DaysOfDevOps-vms-loop-template.json `
   -TemplateParameterFile C:\Users\micha\demo\90DaysOfDevOps\Days\Cloud\02TrafficManagement\Mod06_90DaysOfDevOps-vms-loop-parameters.json

   $location = (Get-AzResourceGroup -ResourceGroupName $rgName).location
   $vmNames = (Get-AzVM -ResourceGroupName $rgName).Name
   
   foreach ($vmName in $vmNames) {
     Set-AzVMExtension `
     -ResourceGroupName $rgName `
     -Location $location `
     -VMName $vmName `
     -Name 'networkWatcherAgent' `
     -Publisher 'Microsoft.Azure.NetworkWatcher' `
     -Type 'NetworkWatcherAgentWindows' `
     -TypeHandlerVersion '1.4'
   }
   