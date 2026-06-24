$rgName = '90DaysOfDevOps'

$webapp = Get-AzWebApp -ResourceGroupName $rgName


#The following following will start an infinite loop that sends the HTTP requests to the web app
while ($true) { Invoke-WebRequest -Uri $webapp.DefaultHostName }