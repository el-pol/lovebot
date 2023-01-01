<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/el-pol/lovebot">
    <img src="img/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Lovebot</h3>

  <p align="center">
    A Twitter bot that tweets love advice every X hours
    <br />
    <a href="https://github.com/el-pol/lovebot"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://twitter.com/loveadvicewiseg">View Demo</a>
    ·
    <a href="https://github.com/el-pol/lovebot/issues">Report Bug</a>
    ·
    <a href="https://github.com/el-pol/lovebot/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Lovebot Screen Shot][lovebot-screenshot]](https://www.polmilian.dev/lovebot_twitter.png)


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Go]][https://go.dev/]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

Install the Go packages by running:
```sh
  go get
```

### Installation

Please follow the steps outlined [here](https://dev.to/elpol/building-a-twitter-bot-with-go-and-gpt-3-57m7) or [here](https://www.polmilian.dev/posts/building-a-twitter-bot-with-go).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Refer to the `Installation` section. You can use this repo as a template for any Twitter bot. The GPT-3 part only relates to fetching the text for our generated tweet; you can replace that part with any valid string.

My bot is live here: [https://twitter.com/LoveAdviceWiseG](https://twitter.com/LoveAdviceWiseG)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] Change the prompts to make it more creative
- [ ] If a user tweets back, reply
- [ ] Put responses in DB & make sure we don't have duplicates

See the [open issues](https://github.com/el-pol/lovebot/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Pol Milian - [@polcodes](https://twitter.com/polcodes) - pol.milian.dev@gmail.com

Project Link: [https://github.com/el-pol/lovebot](https://github.com/el-pol/lovebot)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [AWS Lambda](https://pkg.go.dev/github.com/aws/aws-lambda-go@v1.36.0)
* [Oauth1](https://pkg.go.dev/github.com/dghubble/oauth1@v0.7.2)
* [Godotenv](https://pkg.go.dev/github.com/joho/godotenv@v1.4.0)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/el-pol/lovebot.svg?style=for-the-badge
[contributors-url]: https://github.com/el-pol/lovebot/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/el-pol/lovebot.svg?style=for-the-badge
[forks-url]: https://github.com/el-pol/lovebot/network/members
[stars-shield]: https://img.shields.io/github/stars/el-pol/lovebot.svg?style=for-the-badge
[stars-url]: https://github.com/el-pol/lovebot/stargazers
[issues-shield]: https://img.shields.io/github/issues/el-pol/lovebot.svg?style=for-the-badge
[issues-url]: https://github.com/el-pol/lovebot/issues
[license-shield]: https://img.shields.io/github/license/el-pol/lovebot.svg?style=for-the-badge
[license-url]: https://github.com/el-pol/lovebot/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[Go]:https://go.dev/ 
[linkedin-url]: https://www.linkedin.com/in/pol-milian/
[lovebot-screenshot]: https://www.polmilian.dev/lovebot_twitter.png
