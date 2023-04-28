for /F %%i in ('git describe --long') do ( set commitid=%%i)
set flags="-X github.com/gphper/ginadmin/cmd/cli/version.version=%commitid%"
go build -ldflags %flags% .\cmd\ginadmin