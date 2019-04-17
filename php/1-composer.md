# composer

[composer 中文文档](https://docs.phpcomposer.com/)

> Composer 是 PHP 的一个依赖管理工具。它允许你申明项目所依赖的代码库，它会在你的项目中为你安装他们。

> For Jser: Composer 在PHP中的作用如同 Node 中的 NPM

## 安装 Composer

> 为方便开发请使用全局安装

https://pkg.phpcomposer.com/#how-to-install-composer


## 安装 agent 模块

> agent 模块用于解析浏览器版本信息

在 https://packagist.org/?query=agent 搜索 `agent`

在命令行中运行

```shell
composer require jenssegers/agent
```

安装完成后当前目录会生成 `composer.json` 文件

```json
{
    "require": {
        "jenssegers/agent": "^2.6"
    }
}
```

并且会生成 `vendor/` 文件夹,文件夹中会出现相关依赖.

## 使用 agent 模块

新建 `1-composer/index.php` 文件,文件内容如下:

```html
<?php
// 引用 '../vendor/autoload.php' 文件,就能自动加载 composer.json 中配置的依赖.
require '../vendor/autoload.php';

// use 可以调用命名空间中的类(class)
use Jenssegers\Agent\Agent;

$agent = new Agent();

echo $agent->browser();

```

使用浏览器访问 `/1-composer/index.php` 页面可以看到浏览器版本.

> For Jser: PHP 中的 require 与 JS 中的 require 不同,因为最初 PHP 并没有类似 JS 的 CommonJs 规范所以 PHP require 语法是用来加载文件的,并不是引入文件调用 `module.export` 输出的变量

`use Jenssegers\Agent\Agent;` 可以调用命名空间中的类是因为 `php/vendor/jenssegers/agent/src/Agent.php` 文件中存在如下代码

```html
<?php
namespace Jenssegers\Agent;

use BadMethodCallException;
use Jaybizzle\CrawlerDetect\CrawlerDetect;
use Mobile_Detect;

class Agent extends Mobile_Detect {
// ...
```

`namespace Jenssegers\Agent;` 定义了当前文件的命名空间是 `Jenssegers\Agent`, `class Agent` 定义了当前命名空间下有 `Agent` 类

又因为通过 `require '../vendor/autoload.php';` 自动加载了 `composer.json` 中的依赖,而依赖中存在 `jenssegers/agent`

```js
"require": {
    "jenssegers/agent": "^2.6"
}
```

所以可以通过 `use Jenssegers\Agent\Agent;` 调用 `Agent`;


## 模块化开发

> For Jser: PHP的 `namespace` 是在 PHP5.3.0 中加入的,虽然相对比 Node 的 CommonJs 或 ES-Module 那么自由易用,但是有比没有好.

模块化开发是可维护代码的基础,千万不要所有代码都写在一个php文件,或者通过 `require` 去加载很多 php 文件,而这些文件没有命名空间,就会出现变量污染和冲突的问题.


## psr-4

可以在 `composer.json` 中自己配置 `autoload` psr-4 目录.

```js
"autoload": {
    "psr-0": {
        "App\\": "app/",
    }
}
```

在命令中中运行

```shell
composer dump-autoload
```

> [composer cil ](https://docs.phpcomposer.com/03-cli.html)
