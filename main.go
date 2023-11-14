package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type planet struct {
	ID					  string `json:"id"`
	Name				  string `json:"name"`
	Description 		  string `json:"description"`
	Mass                  string `json:"mass"`
	Diameter              string `json:"diameter"`
	OrbitPeriod           string `json:"orbitPeriod"`
	DistanceFromSun       string `json:"distanceFromSun"`
	NumberOfMoons         int    `json:"numberOfMoons"`
	HottestTemperature    int `json:"hottestTemperature"`
	ColdestTemperature    int `json:"coldestTemperature"`
	PictureUrl			  string `json:"pictureUrl"`
	ThreeDModelUrlDark	  string `json:"threeDModelUrlDark"`
	ThreeDModelUrlLight	  string `json:"threeDModelUrlLight"`
	InterestingFact		  string `json:"interestingFact"`
}

var planets = []planet {
	{
		ID: "0",
		Name: "Sun",
		Description: "The Sun, at the center of our solar system, is a hot ball of glowing gases, providing the heat and light that sustain life on Earth.",
		Mass: "1.989 × 10^30",
		Diameter: "1,391,000",
		OrbitPeriod: "N/A",
		DistanceFromSun: "0",
		NumberOfMoons: 0,
		HottestTemperature: 27000000,
		ColdestTemperature: 0,
		PictureUrl: "https://i.imgur.com/UP3jh7e.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/sun/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/sun/desktop_light.mp4",
		InterestingFact: "The Sun contains 99.86% of the mass in the Solar System and is about 109 times the diameter of Earth.",
	},
	{
		ID: "1",
		Name: "Mercury",
		Description: "Mercury, the smallest and closest planet to the Sun, is known for its cratered surface, lack of moons, and extreme temperature variations.",
		Mass: "3.285 × 10^23",
		Diameter: "3,031.9",
		OrbitPeriod: "88 days",
		DistanceFromSun: "43.245 million miles",
		NumberOfMoons: 0,
		HottestTemperature: 800,
		ColdestTemperature: -290,
		PictureUrl: "https://i.imgur.com/jTJbpj0.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/mercury/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/mercury/desktop_light.mp4",
		InterestingFact: "Not only is Mercury the smallest planet, it is also shrinking!",
	},
	{
		ID: "2",
		Name: "Venus",
		Description: "Venus, often called Earth's sister planet, is shrouded in thick clouds, experiences extreme greenhouse effects, and has the hottest surface temperatures in the solar system.",
		Mass: "4.867 × 10^24",
		Diameter: "7,520.8",
		OrbitPeriod: "225 days",
		DistanceFromSun: "66.843 million miles",
		NumberOfMoons: 0,
		HottestTemperature: 900,
		ColdestTemperature: -226,
		PictureUrl: "https://i.imgur.com/MtBAFmy.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/venus_surface/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/venus_surface/desktop_light.mp4",
		InterestingFact: "It takes Venus longer to rotate once on its axis than to complete one orbit of the Sun!",
	},
	{
		ID: "3",
		Name: "Earth",
		Description: "Earth, the only planet known to harbor life, has diverse ecosystems and a unique atmosphere that sustains life.",
		Mass: "5.972 × 10^24",
		Diameter: "7,917.5",
		OrbitPeriod: "365 days",
		DistanceFromSun: "92.96 million miles",
		NumberOfMoons: 1,
		HottestTemperature: 134,
		ColdestTemperature: -128,
		PictureUrl: "https://i.imgur.com/G3YZQEU.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/earth/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/earth/desktop_light.mp4",
		InterestingFact: "Earth is the only planet not named after a mythological god or goddess!",
	},
	{
		ID: "4",
		Name: "Mars",
		Description: "Mars, known as the Red Planet, is famous for its red soil, thin atmosphere, and potential for past water presence.",
		Mass: "6.417 × 10^23",
		Diameter: "4,212.3",
		OrbitPeriod: "687 days",
		DistanceFromSun: "141.6 million miles",
		NumberOfMoons: 2,
		HottestTemperature: 70,
		ColdestTemperature: -225,
		PictureUrl: "https://i.imgur.com/bR5orPA.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/mars/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/mars/desktop_light.mp4",
		InterestingFact: "Mars has the tallest volcano in the solar system, Olympus Mons.",
	},
	{
		ID: "5",
		Name: "Jupiter",
		Description: "Jupiter, the largest planet in the solar system, is a gas giant known for its Great Red Spot and numerous moons.",
		Mass: "1.898 × 10^27",
		Diameter: "86,881",
		OrbitPeriod: "12 years",
		DistanceFromSun: "484.3 million miles",
		NumberOfMoons: 95,
		HottestTemperature: 43000,
		ColdestTemperature: -260,
		PictureUrl: "https://i.imgur.com/FpI4wEi.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/jupiter/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/jupiter/desktop_light.mp4",
		InterestingFact: "Jupiter has the shortest day of all the planets, rotating once in just under 10 hours.",
	},
	{
		ID: "6",
		Name: "Saturn",
		Description: "Saturn is renowned for its spectacular ring system and is a gas giant like Jupiter.",
		Mass: "5.683 × 10^26",
		Diameter: "74,898",
		OrbitPeriod: "29 years",
		DistanceFromSun: "886.7 million miles",
		NumberOfMoons: 146,
		HottestTemperature: 21000,
		ColdestTemperature: -191,
		PictureUrl: "https://i.imgur.com/tyErvuB.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/saturn/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/saturn/desktop_light.mp4",
		InterestingFact: "Saturn is the only planet in the solar system that is less dense than water.",
	},
	{
		ID: "7",
		Name: "Uranus",
		Description: "Uranus, an ice giant, is unique for its tilted rotation axis and has a pale blue color due to methane in its atmosphere.",
		Mass: "8.681 × 10^25",
		Diameter: "31,518",
		OrbitPeriod: "84 years",
		DistanceFromSun: "1,787 billion miles",
		NumberOfMoons: 27,
		HottestTemperature: -9000,
		ColdestTemperature: -224,
		PictureUrl: "https://i.imgur.com/Hlqh4Zz.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/uranus/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/uranus/desktop_light.mp4",
		InterestingFact: "Uranus rotates on its side, making its seasons last 21 years each.",
	},
	{
		ID: "8",
		Name: "Neptune",
		Description: "Neptune, the farthest known planet from the Sun, is an ice giant with strong winds and a deep blue color.",
		Mass: "1.024 × 10^26",
		Diameter: "49,244",
		OrbitPeriod: "165 years",
		DistanceFromSun: "2,798",
		NumberOfMoons: 14,
		HottestTemperature: 12632,
		ColdestTemperature: -214,
		PictureUrl: "https://i.imgur.com/D7U1LJc.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/neptune/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/neptune/desktop_light.mp4",
		InterestingFact: "Neptune was the first planet located through mathematical predictions rather than through regular observation.",
	},
	{
		ID: "9",
		Name: "Pluto",
		Description: "Pluto, a dwarf planet in the Kuiper belt, is known for its icy surface and was the ninth planet until reclassified in 2006.",
		Mass: "1.309 × 10^22",
		Diameter: "1,476.8",
		OrbitPeriod: "248 years",
		DistanceFromSun: "5,906",
		NumberOfMoons: 5,
		HottestTemperature: -232,
		ColdestTemperature: -387,
		PictureUrl: "https://i.imgur.com/6JgWf9A.png",
		ThreeDModelUrlDark: "https://www.gstatic.com/culturalinstitute/searchar/assets/pluto/desktop_dark.mp4",
		ThreeDModelUrlLight: "https://www.gstatic.com/culturalinstitute/searchar/assets/pluto/desktop_light.mp4",
		InterestingFact: "Pluto has a heart-shaped glacier the size of Texas and Oklahoma combined.",
	},
}

func CORSMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	}
}

func getPlanets(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, planets)
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/planets", getPlanets)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}