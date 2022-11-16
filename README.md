# Discord bots

## Change status bot.

```go
func main() {
	token := "YOUR_BOT_TOKEN"
	newBot := discord.New(token, "YOUR_BOT_NAME", time.Second*2)
	err := newBot.Open()
	if err != nil {
		panic(err)
	}
	newBot.ChangeNames("NAME", "YOUR_GUILD_ID")
	newBot.ChangePlaying("PLAYING")
}
```

## Reply bot.

```go
func main() {
	token := "YOUR_BOT_TOKEN"
	newBot := discord.New(token, "YOUR_BOT_NAME", time.Second*2)
	err := newBot.Open()
	if err != nil {
		panic(err)
	}
	newBot.AddReplyRule(discord.ReplyRule{
		ChannelID: "YOUR_CHANNEL_ID",
		RuleType:  0, // Equalfold.
		CheckText: "!equalfold",
		ReplyText: "Reply",
	})
	newBot.AddReplyRule(discord.ReplyRule{
		ChannelID: "1006871784143999016",
		RuleType:  1,// Contain.
		CheckText: "!contain",
		ReplyText: "Reply",
	})
	newBot.StartReply()
}
```
