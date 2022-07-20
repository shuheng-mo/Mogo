<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!--
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- using the static badge because it is private, covert to dynamic ones if public  -->
<!-- https://shields.io/#your-badge -->
<div>
<img src="https://img.shields.io/badge/test-failed-red">
<!-- <img src="https://img.shields.io/badge/test-passing-green"> -->
<img src="https://img.shields.io/badge/issues-1-red">
<img src="https://img.shields.io/badge/contributors-1-green">
<img src="https://img.shields.io/badge/license-MIT-yellow">
</div>

<!-- PROJECT LOGO -->
<div align="center">
  <!-- <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h1 align="center">MOGO</h1>
  <p align="center">
    <a href="https://github.com/ese-msc-2021/irp-sm321/tree/main/documentation"><strong>Read the code documentation »</strong></a>
    <br />
    <br />
    <a href="https://github.com/ese-msc-2021/irp-sm321">View Demo</a>
    ·
    <a href="https://github.com/ese-msc-2021/irp-sm321/issues">Report Bug</a>
    ·
    <a href="https://github.com/ese-msc-2021/irp-sm321/issues">Request Feature</a>
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

A micro-service application written by Go.

Core features:
*  RESTful API

Project structure:
```
```

<p align="right">(<a href="#top">back to top</a>)</p>

### Code Metadata

This section listed the major frameworks/libraries used to bootstrap this project. Other add-ons/plugins please refer to the acknowledgements section.

* [go-micro](https://mpi4py.readthedocs.io/en/stable/)


<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

Here is an example of how you may give instructions on setting up this project locally. To get a local copy up and running follow these simple example steps.

### Prerequisites

There are some prerequisites before you compile and run the project on local machine or your AI computers. Note that this project built by Python language and relevant packages, add-ons, dependencies.
* Python [Installation guide](https://www.python.org/downloads/)
* Anaconda (optional) [Installation guide](https://docs.conda.io/projects/conda/en/latest/user-guide/install/index.html)
* mpi4py [Installation guide](https://mpi4py.readthedocs.io/en/latest/install.html)
  ```bash
  # using pip
  $ python -m pip install mpi4py
  # using conda (mpich, openmpi recommended on HPC)
  $ conda install -c conda-forge mpi4py mpich
  ```

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo
   ```sh
   git clone https://github.com/ese-msc-2021/irp-sm321
   ```
2. Install required packages / compile
    ```
    pip install -e .
    ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

Code example:
```
# runing the parallel scripts
$ mpiexec -n 16 python parallel_advection_diffusion_2D.py
```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] The first objective.

See the [open issues](https://github.com/ese-msc-2021/irp-sm321/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the project such an amazing thing to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/your_feature`)
3. Commit your Changes (`git commit -m 'Add some new feature'`)
4. Push to the Branch (`git push origin feature/your_feature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See [`LICENSE.md`](https://github.com/ese-msc-2021/irp-sm321/blob/main/LICENSE.md) for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

## References
- [Some book title here]()

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->
## Acknowledgments
* [me](sm321@ic.ac.uk)


<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Shuheng Mo - [sm321@ic.ac.uk](shuheng.mo21@imperial.ac.uk)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png

<!-- ## About IRP project

We generated this repository for you to work in during your IRP. Please note that if your code is confidential, you do not need to submit your code here.

- For project information updates, please refer to `README.md` in [`info/`](./info) directory.
- For details on deliverable submissions, please refer to `README.md` in [`reports/`](./reports) directory.

If you have any questions or experience any difficulties, please do not hesitate to get in touch with Marijan (m.beg@imperial.ac.uk).

Please feel free to remove this text from `README.md` and write it the way you want. ;) -->

