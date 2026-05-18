package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protocol_buffers/protobuff"
)

func main() {
	fmt.Println("hello world")
	socials := protobuff.SocialFollowers{
		Youtube: 1400,
		Twitter: 2500,
	}
	elliot := &protobuff.Person{
		Name:            "Elliot",
		Age:             18,
		Socialfollowers: &socials,
	}
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshalling error: ", err)
	}

	fmt.Println(data)

	// reterival
	newElliot := &protobuff.Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}

	fmt.Println(newElliot)
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetSocialfollowers())
	fmt.Println(newElliot.GetSocialfollowers().GetTwitter())
	fmt.Println(newElliot.GetSocialfollowers().GetYoutube())
}
