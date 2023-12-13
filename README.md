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
[![GPL-3.0 License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/ofman/fsgui">
    <img src="img/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">fsgui (filesharego-ui)</h3>

  <p align="center">
    File sharing P2P gui+cli program made in GoLang with IPFS and Fyne via KISS (Keep It Super Simple) principle. 
    <br />
    <a href="https://github.com/ofman/fsgui"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <!-- <a href="https://github.com/ofman/fsgui">View Online Demo</a>
    · -->
    <a href="https://github.com/ofman/fsgui/issues">Report Bug</a>
    ·
    <a href="https://github.com/ofman/fsgui/issues">Request Feature</a>
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

[![FileShareGo-cli Screen Shot][product-screenshot]](img/screenshot.png)

There are many file sharing programs on GitHub; however, I didn't find one that really suited my needs so I created this enhanced one. I want to create a program so amazing that it'll be the last one you ever need -- I think this is it.

Here's why:
* Your time should be focused on creating something amazing. A project that solves a problem and helps others
* You shouldn't be doing the same tasks over and over like zipping/extracting or slicing files in parts just to send them
* You should implement KISS principles to the rest of your life :smile:

Of course, no one program will serve all projects since your needs may be different. So I'll be adding more in the near future (for example filesharego-gui). You may also suggest changes by forking this repo and creating a pull request or opening an issue. Thanks to all the people have contributed to expanding this program!

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][go.dev]][go-url]
* [![IPFS][ipfs.tech]][ipfs-url]
* [![Debian][debian.org]][debian-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Only need 1 binary file "fsg" to run which can be quickly downloaded from releases section: https://github.com/ofman/fsgui/releases/latest

To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.


* ubuntu/debian
  ```sh
  apt-get -y install wget
  wget https://github.com/ofman/fsgui/releases/download/release-v0.0.1/fsg
  cp fsg /usr/local/bin
  ```

### Installation

1. Install git
   ```sh
   apt-get -y install git
   ```
2. Clone the repo
   ```sh
   git clone https://github.com/ofman/fsgui.git
   ```
3. Install GO from official website here https://go.dev/doc/install or bellow via apt and check version
   ```sh
   apt-get install golang
   export PATH=$PATH:/usr/local/go/bin
   go version
   ```
4. cd into folder and build
   ```sh
   cd fsg && go mod tidy && go build
   ```
5. copy binary file to /usr/local/bin
   ```sh
   cp fsg /usr/local/bin
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use flags -f "example.jpg" or -c "exampleCid" to share files for example:
Upload file (keep terminal window open/running to let others download):
   ```sh
   ./fsgui -f example.jpg
   ```
Download file (open new terminal window):
   ```sh
   ./fsgui -c /ipfs/QmX4zdEUtimXgxhpzv8jfFLqkuutNhmoNH987cH5RS67GM
   ```

_For more examples, please refer to the [Documentation](https://oofman.github.com/fsg-site)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] Single file download in the same directory
- [ ] Website
- [ ] Public seed server helpers
    - [ ] Blockchain seed/leech ratio

See the [open issues](https://github.com/ofman/fsgui/issues) for a full list of proposed features (and known issues).

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

Distributed under the GNU-V3 License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

ofman<!-- [@twitter_handle](https://twitter.com/twitter_handle) -->- ofmangit@gmail.com

Project Link: [https://github.com/ofman/fsgui](https://github.com/ofman/fsgui)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* []()
* []()
* [![Debian][debian.org]][debian-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/ofman/fsg.svg?style=for-the-badge
[contributors-url]: https://github.com/ofman/fsgui/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ofman/fsg.svg?style=for-the-badge
[forks-url]: https://github.com/ofman/fsgui/network/members
[stars-shield]: https://img.shields.io/github/stars/ofman/fsg.svg?style=for-the-badge
[stars-url]: https://github.com/ofman/fsgui/stargazers
[issues-shield]: https://img.shields.io/github/issues/ofman/fsg.svg?style=for-the-badge
[issues-url]: https://github.com/ofman/fsgui/issues
[license-shield]: https://img.shields.io/github/license/ofman/fsg.svg?style=for-the-badge
[license-url]: https://github.com/ofman/fsgui/blob/master/LICENSE
[product-screenshot]: img/screenshot.png
[go.dev]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[go-url]: https://go.dev/
[ipfs.tech]: https://img.shields.io/badge/IPFS-Inter_Planetary_File_System-blue
[ipfs-url]: https://ipfs.tech/
[debian.org]: https://img.shields.io/badge/Debian-A81D33?style=for-the-badge&logo=debian&logoColor=white
[debian-url]: https://www.debian.org/distrib/