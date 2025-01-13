---
title: Home
layout: home
nav_order: 1
description: "Documentation for Bloader, the modern benchmarking tool that simplifies load testing. Whether you're a seasoned developer or just starting, Bloader provides flexibility and power for all your testing needs"
permalink: /
---

<h1 align="center">
  <a href="https://docs.bloader.cresplanex.org">
    <picture>
      <source height="60" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/cresplanex/bloader/main/docs/static/bloader_logo.png">
      <img height="60" alt="Bloader" src="https://raw.githubusercontent.com/cresplanex/bloader/main/docs/static/bloader_logo.png">
    </picture>
  </a>
  <br>
</h1>

Bloader is a benchmark testing project that focuses on flexibility and simplicity.
{: .fs-6 .fw-300 }

[Get started now](setup/index.md){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 }
[View it on GitHub][Bloader repo]{: .btn .fs-5 .mb-4 .mb-md-0 }

---

{: .warning }
> This website documents the features of the current `main` branch of Bloader. See [the CHANGELOG]({% link CHANGELOG.md %}) for a list of releases, new features, and bug fixes.

Welcome to the official documentation for **Bloader**, the modern benchmarking tool that simplifies load testing. Whether you're a seasoned developer or just starting, Bloader provides flexibility and power for all your testing needs.

### 🛠 Features
- Internal Store for managing requests.
- **Master-Slave Architecture** with gRPC for distributed testing.
- YAML-based configuration with **Sprig** template engine.

Browse the docs to learn more about how to use bloader.

### 📖 Reference Software
- [Sprig](https://masterminds.github.io/sprig/): Template engine 
- [Cobra](https://github.com/spf13/cobra): CLI creation 
- [Viper](https://github.com/spf13/viper): Parsing Config files
- [Buf](https://buf.build/): gRPC schema management 
- [Bolt](https://github.com/boltdb/bolt): Internal store 

## Future implementations 
- Change from BoltDB to a database that is still supported today
- Add external cloud providers, etc. to Override's Type.
- Add functionality for performing analysis.
- Make it run as a server and make it visually clear besides the CLI.
- Introduce original Encrypt between Master and Slave.
- Add gRPC as a measurement target.
- Addition of test code 
- Add plugin functionality. 

## About the project

Bloader is &copy; 2024-{{ "now" | date: "%Y" }} by [Cresplanex](open-source-github@cresplanex.com).

### License

Bloader is distributed by an [MIT license](https://github.com/cresplanex/bloader/tree/main/LICENSE).

### Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change. Read more about becoming a contributor in [our GitHub repo](https://github.com/cresplanex/bloader/blob/main/docs/contributing/index.md).

#### Thank you to the contributors of Bloader!

<ul class="list-style-none">
{% for contributor in site.github.contributors %}
  <li class="d-inline-block mr-1">
     <a href="{{ contributor.html_url }}"><img src="{{ contributor.avatar_url }}" width="32" height="32" alt="{{ contributor.login }}"></a>
  </li>
{% endfor %}
</ul>

### Code of Conduct

Bloader is committed to fostering a welcoming community.

[View our Code of Conduct](https://github.com/cresplanex/bloader/tree/main/.github/CODE_OF_CONDUCT.md) on our GitHub repository.

[Bloader]: https://docs.bloader.cresplanex.org
[Bloader repo]: https://github.com/cresplanex/bloader
[Bloader README]: https://github.com/cresplanex/bloader/blob/main/README.md
[GitHub Pages]: https://pages.github.com/

