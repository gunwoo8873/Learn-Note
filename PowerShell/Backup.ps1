$date =
Get-Date -format "yyyy-MM-dd"
Compress-Archive -Path '.' -CompressionLevel 'Fastest' -DestinationPath "./backup-$date"
Write-Host "Created backup at $('./backup-' + $date + '.zip')"