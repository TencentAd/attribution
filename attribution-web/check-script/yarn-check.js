const script = process.env.npm_lifecycle_event;
const excutePath = process.env.npm_execpath;
const isYarn = excutePath.endsWith("yarn.js");

if (!isYarn) {
  console.log(`\x1b[31m`);
  console.log("请使用 yarn");
  console.log("项目强制使用锁版本");
  console.info("更多信息可以查看 yarn 官网: https://yarnpkg.com/lang/en/");
  process.exit(-1);
}
