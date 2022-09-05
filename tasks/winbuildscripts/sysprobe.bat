if not exist c:\mnt\ goto nomntdir

@echo c:\mnt found, continuing

@echo PARAMS %*
@echo PY_RUNTIMES %PY_RUNTIMES%

if NOT DEFINED PY_RUNTIMES set PY_RUNTIMES=%~1

call %~p0extract-modcache.bat

set BUILD_ROOT=c:\buildroot
mkdir %BUILD_ROOT%\datadog-agent
if not exist %BUILD_ROOT%\datadog-agent exit /b 2
cd %BUILD_ROOT%\datadog-agent || exit /b 3
xcopy /e/s/h/q c:\mnt\*.* || exit /b 4


Powershell -C "c:\mnt\tasks\winbuildscripts\sysprobe.ps1" || exit /b 5

REM copy resulting packages to expected location for collection by gitlab.
if not exist c:\mnt\test\kitchen\site-cookbooks\dd-system-probe-check\files\default\tests\ mkdir c:\mnt\test\kitchen\site-cookbooks\dd-system-probe-check\files\default\tests\ || exit /b 6
xcopy /e/s/q %BUILD_ROOT%\datadog-agent\test\kitchen\site-cookbooks\dd-system-probe-check\files\default\tests\*.* c:\mnt\test\kitchen\site-cookbooks\dd-system-probe-check\files\default\tests\ || exit /b 7
