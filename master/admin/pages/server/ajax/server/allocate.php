<?php
/*
    PufferPanel - A Minecraft Server Management Panel
    Copyright (c) 2013 Dane Everitt
 
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.
 
    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.
 
    You should have received a copy of the GNU General Public License
    along with this program.  If not, see http://www.gnu.org/licenses/.
 */
session_start();
require_once('../../../../../core/framework/framework.core.php');

if($core->framework->auth->isLoggedIn($_SERVER['REMOTE_ADDR'], $core->framework->auth->getCookie('pp_auth_token'), true) !== true){
	$core->framework->page->redirect('../../../../index.php');
}

if(!isset($_POST['sid']))
	$core->framework->page->redirect('../../find.php');

/*
 * Validate Disk & Memory
 */	
if(!is_numeric($_POST['alloc_mem']) || !is_numeric($_POST['alloc_disk']))
	$core->framework->page->redirect('../../view.php?id='.$_POST['sid'].'&error=alloc_mem|alloc_disk&disp=m_fail');

$mysql->prepare("UPDATE `servers` SET `max_ram` = :ram, `disk_space` = :disk WHERE `id` = :sid")->execute(array(
    ':sid' => $_POST['sid'],
    ':ram' => $_POST['alloc_mem'],
    ':disk' => $_POST['alloc_disk']
));

$core->framework->page->redirect('../../view.php?id='.$_POST['sid']);