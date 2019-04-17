<?php
// 引用 'vendor/autoload.php' 文件,就能自动加载 composer.json 中配置的依赖.
require '../vendor/autoload.php';

// use 可以调用命名空间中的类(class)
use Jenssegers\Agent\Agent;

$agent = new Agent();

echo $agent->browser();
