<?php
setcookie("cookie_name", "cookie_value", time()+60*60*24*30, "/", "", false, true);
echo $_COOKIE["cookie_name"];
