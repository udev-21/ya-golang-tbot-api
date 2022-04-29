## Telegram Passport Example

### Generating a private key
First, use a console to generate a private key:
#### WARNING: Keep your private key SECRET! 
```sh
openssl genrsa 2048 > private.key
```
### Generating your public key
Then use the console to print the corresponding public key:
```sh
openssl rsa -in private.key -pubout
```

Use the /setpublickey command with [@BotFather](https://t.me/BotFather) to connect this public key with your bot.

### And then...
Configure your index.html with your 
```html
<div id="telegram_passport_auth"></div>
<script>
    Telegram.Passport.createAuthButton('telegram_passport_auth', {
        bot_id:       11223344, // place id of your bot here
        // abouto scope here: https://core.telegram.org/passport/#passportscope
        scope:        {data: [{type: 'id_document', selfie: true}, 'address_document', 'email'], v: 1}, 
        public_key:   '-----BEGIN PUBLIC KEY-----...', // place public key of your bot here
        nonce:        'ab2df83746a87d2f3bd6', // place nonce here 
    });
</script>
```

after that open index.html on your browser and click on button, after 
telegram successfully send your data to bot, run example.go file:
```sh
go run example.go
```
