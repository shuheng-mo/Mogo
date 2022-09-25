<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- using the static badge because it is private, covert to dynamic ones if public  -->
<!-- https://shields.io/#your-badge -->

<div>
<img src="https://img.shields.io/github/issues/shuheng-mo/Mogo">
<img src="https://img.shields.io/github/forks/shuheng-mo/Mogo">
<img src="https://img.shields.io/github/stars/shuheng-mo/Mogo">
<img src="https://img.shields.io/github/license/shuheng-mo/Mogo">
</div>

<!-- PROJECT LOGO -->
<div align="center">
  <!-- <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h1 align="center">MOGO</h1>
  <p align="center">
    <a href="https://github.com/shuheng-mo/Mogo"><strong>Read the code documentation »</strong></a>
    <br />
    <br />
    <a href="https://github.com/shuheng-mo/Mogo">View Demo</a>
    ·
    <a href="https://github.com/shuheng-mo/Mogo/issues">Report Bug</a>
    ·
    <a href="https://github.com/shuheng-mo/Mogo/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#code-metadata">Code Metadata</a></li>
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
    <li><a href="#references">References</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

[![Test](https://github.com/shuheng-mo/Mogo/workflows/Test/badge.svg)](https://github.com/shuheng-mo/Mogo/actions)


A micro-service E-commerce application written by Go.

Core features:
*  RESTful API
*  Paypal payment and refund

This project followed the project structure provided by [micro](https://hub.docker.com/r/micro/micro), which open-sourced and licensed.
```bash
$ docker pull micro/micro
```

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>

### Code Metadata

This section listed the major frameworks/libraries used to bootstrap this project. Other add-ons/plugins please refer to the acknowledgements section.

* go-micro
* gRPC
* protobuf (version 3)
* gorm
* [Consul](https://github.com/hashicorp/consul)
* [Jaeger](https://github.com/jaegertracing/jaeger)
* [hystrix-go](https://github.com/afex/hystrix-go)
* [grafana](https://grafana.com/grafana/dashboards/)
* [Prometheus](https://prometheus.io/docs/guides/go-application/)
* [ELK stack](https://www.elastic.co/elastic-stack?ultron=B-Stack-Trials-EMEA-UK-Exact&gambit=Stack-ELK&blade=adwords-s&hulk=paid&Device=c&thor=elk%20stack%20logging&gclid=Cj0KCQjwguGYBhDRARIsAHgRm49c4_7YcuZLwIs3Hi1MK0R4WAWhV4EuQ4PEzSh-dBROKHCt80AyxfYaAuxYEALw_wcB)
* [zap](https://github.com/uber-go/zap)
* Docker
* K8S

Each micro-service module followed the project structure provided by [micro](https://hub.docker.com/r/micro/micro).

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

Here is an example of how you may give instructions on setting up this project locally. To get a local copy up and running follow these simple example steps.

### Prerequisites

There are some prerequisites before you compile and run the project on local machine or your AI computers. Note that this project built by Python language and relevant packages, add-ons, dependencies.
* [Go](https://go.dev/)
* MySQL
* Docker

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo
   ```sh
   git clone https://github.com/shuheng-mo/Mogo
   ```
2. Install customized wheels
    ```
    $ pip install -e .
    # if use setup.py use pip install .
    ```
3. Set up go modules
    ```bash
      # update go dependencies
      $ go mod tidy
      # update required dependencies
      $ go get ...
    ```
4. Set up new `key-value` pairs in Consul like:
    ```
    {
      "host":"127.0.0.1",
      "user":"root",
      "pwd":"123456",
      "database":"micro",
      "port":3306
    }
    ```

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

Code example:
```
# Executing
$ ./app

$ 
```

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Install MySQL and create databases (tables).
- [x] DDD model template for each module.

See the [open issues](https://github.com/shuheng-mo/Mogo/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the project such an amazing thing to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/your_feature`)
3. Commit your Changes (`git commit -m 'Add some new feature'`)
4. Push to the Branch (`git push origin feature/your_feature`)
5. Open a Pull Request

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>



<!-- LICENSE -->
## License

Distributed under the GPL-3.0 License. See [`LICENSE.md`](https://github.com/shuheng-mo/Mogo/blob/main/LICENSE) for more information.

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>

## References
- Docker containers 
  - used docker containers ...

- Github repositories

- Book and publications

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments
This project won't born without the help of these wonderful people/coporations:
*

<p align="right">(<a href="#top">BACK TO TOP</a>)</p>

<!-- CONTACT -->
## Contact

Shuheng Mo - [Contact me](https://linktr.ee/shuheng_mo)


<p align="right">(<a href="#top">BACK TO TOP</a>)</p>



