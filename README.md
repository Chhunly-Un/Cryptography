# Cryptography
# Un Chhunly

# How to run code

# For Lab# and Lab#4
cd ".\Un Chhunly(G02) Assignment week3"
go run .\Lab#0_proof_hash.go
go run .\Lab#4-Mini_CTF.go

# For Lab1 don't need to input anything just run and will see result
cd ".\Un Chhunly(G02) Assignment week3\lab1_md5_cracker\"
go run .\main.go

# For Lab2 need to input sha1 after running main.go will see result
sha1: aa1c7d931cf140bb35a5a16adeb83a551649c3b9
cd ".\Un Chhunly(G02) Assignment week3\Lab2_password_cracker\"
go run .\main.go

# For Lab3 need to input sha512 after running main.go will see result
Sha512:485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38
cd ".\Un Chhunly(G02) Assignment week3\Lab3_password-sha512_cracker\"
go run .\main.go

# Push to github
git init
git add .
git commit -m "Initial commit"
git remote add origin https://github.com/Chhunly-Un/Cryptography.git
git branch -M main
git pull origin main --allow-unrelated-histories
git push -u origin main