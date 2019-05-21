$filePath = Get-ChildItem -Path (Read-Host -Prompt 'Enter path to file')
foreach($line in Get-Content $filePath) {
    $sha256 = new-object -TypeName System.Security.Cryptography.SHA256CryptoServiceProvider
    $utf8 = new-object -TypeName System.Text.UTF8Encoding
    $hash = [System.BitConverter]::ToString($sha256.ComputeHash($utf8.GetBytes($line)))
    $hashWithoutSeparator = ($hash) -replace "-", ""

    Write-Output $hashWithoutSeparator
}
