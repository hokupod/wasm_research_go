<?php
setcookie("cookie_name", "cookie_value", time()+60*60*24*30, "/", $_GET["domain"], false, true);
var_dump($_GET["domain"]);
