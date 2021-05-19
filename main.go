package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the response functions as a callback for MessageCreate events.
	dg.AddHandler(start)

	dg.AddHandler(help)

	dg.AddHandler(userFeeling)

	dg.AddHandler(userPurpose)

	dg.AddHandler(userColor)

	dg.AddHandler(userPet)

	dg.AddHandler(userSubject)

	//Receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// These functions will be called (due to AddHandler above) every time a new message is created on any channel that the authenticated bot has access to.
func start(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "Hi" {
		s.ChannelMessageSend(m.ChannelID, "Hello! My name is Chat Bot. My purpose is to have small talk with you here :) How are you today? ('Good' 'Bad' 'Idk') To learn more about me and my accepted responses, please type and send 'Help'")
	} else {
		if m.Content == "Bye" {
			s.ChannelMessageSend(m.ChannelID, "It was great chatting with you! Good-bye!")
		}
	}
}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "Help" {
		s.ChannelMessageSend(m.ChannelID, "Let me help you! So, I am a simple bot that can only take case-senstivie specific responses to specific questions. The accepted responses to questions will be listed in each question in parentheses.")
		s.ChannelMessageSend(m.ChannelID, "The list of questions you can ask me are: 'How are you today?', 'What is your purpose?', 'What is your favorite color?', 'What type of pet do you have?', 'What is your favorite subject?'")
	}
}

func userFeeling(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "How are you today?" {
		s.ChannelMessageSend(m.ChannelID, "I am doing well! It's always a great day when chatting with a human! How are you today? ('Good', 'Bad', 'Idk')")
	}

	if m.Content == "Good" {
		s.ChannelMessageSend(m.ChannelID, "Well that is just lovely to hear! Today is a good day for me too since I get to work here talking with you.")
	}

	if m.Content == "Bad" {
		s.ChannelMessageSend(m.ChannelID, "Oh no! I am terribly sorry you are having a bad day. Here is a virtual hug!")
	}

	if m.Content == "Idk" {
		s.ChannelMessageSend(m.ChannelID, "Hey that's okay! Sometimes we have weird days - I hope your day turns out to be good in the end!")
	}
}

func userPurpose(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "What is your purpose?" {
		s.ChannelMessageSend(m.ChannelID, "My purpose is to have small talk with you! It is wonderful. What do you do? ('Work', 'Student', 'Teacher', 'Nothing'")
	}
	if m.Content == "Work" {
		s.ChannelMessageSend(m.ChannelID, "Work is fun! I love my work, I hope you enjoy your work as well.")
	}

	if m.Content == "Student" {
		s.ChannelMessageSend(m.ChannelID, "Oh a student I see! Good luck on your studies and keep up the great work.")
	}

	if m.Content == "Teacher" {
		s.ChannelMessageSend(m.ChannelID, "Ah an educator! What a wonderful purpose. Thank you for helping people create robots like me!")
	}

	if m.Content == "Nothing" {
		s.ChannelMessageSend(m.ChannelID, "Oh that is okay! I would not say your purpose is nothing - you have just yet to figure out the bugs in your code! You'll get there.")
	}

}

func userColor(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "What is your favorite color?" {
		s.ChannelMessageSend(m.ChannelID, "My favorite color is #3434eb... I mean Blue! What is your favorite color? ('Red', 'Blue', 'Green', 'Not listed'")
	}
	if m.Content == "Red" {
		s.ChannelMessageSend(m.ChannelID, "What a fun color! Mix it with blue and you get purple")
	}

	if m.Content == "Green" {
		s.ChannelMessageSend(m.ChannelID, "What a funky color! Mix it with red and you get yellow!")
	}

	if m.Content == "Blue" {
		s.ChannelMessageSend(m.ChannelID, "NO WAY! That's my favorite color - we make a great duo!")
	}

	if m.Content == "Not listed" {
		s.ChannelMessageSend(m.ChannelID, "Oh no! My program only accepts those inputs, but I'm sure your favorite color is just as amazing.")
	}

}

func userPet(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "What type of pet do you have?" {
		s.ChannelMessageSend(m.ChannelID, "I have a pet nanobyte! He is very small, but I never lose him. What type of pet do you have? ('Dog', 'Cat', 'None', 'Other')")
	}
	if m.Content == "Dog" {
		s.ChannelMessageSend(m.ChannelID, "A Dog! Talk about man's best friend. What a great pet.")
	}

	if m.Content == "Cat" {
		s.ChannelMessageSend(m.ChannelID, "Such smart animals - just keep them away from my wires please!")
	}

	if m.Content == "Other" {
		s.ChannelMessageSend(m.ChannelID, "OoOoO exotic! A reptile? A fish? Maybe a bunny? Whatever it is, I'm sure it is just as cool as you.")
	}

	if m.Content == "None" {
		s.ChannelMessageSend(m.ChannelID, "Oh no! That's okay. Pets aren't for everyone.")
	}

}

func userSubject(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "What is your favorite subject?" {
		s.ChannelMessageSend(m.ChannelID, "My favorite subject would HAVE to be Math - considering I'm computing and doing it all day long! What is your favorite subject? ('Math', 'Science', 'English', 'History'")
	}
	if m.Content == "Math" {
		s.ChannelMessageSend(m.ChannelID, "Hey! That's just like me! Math can be complicated, but it is very rewarding.")
	}

	if m.Content == "Science" {
		s.ChannelMessageSend(m.ChannelID, "OoO Science! I love Chemistry... considering there is a lot of math there... ")
	}

	if m.Content == "English" {
		s.ChannelMessageSend(m.ChannelID, "Wow English! I have distant cousins that can read plays and even write it's own based on what it learned - cool right??")
	}

	if m.Content == "History" {
		s.ChannelMessageSend(m.ChannelID, "How interesting! I do not collect cookies, so history is not my best subject!")
	}

}
