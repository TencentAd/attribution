# cd 到 attribution-web 目录
CURRENT_DIR=$(cd $(dirname $0); pwd)
cd ${CURRENT_DIR}
cd ../

yarn && yarn build

rm -rf ./node_modules

# 创建 cdn 临时文件夹
mkdir -p ./cdn-release/files/attribution-web

# 将 deploy 文件夹及多个静态文件夹迁移到 cdn-release 文件夹下
rsync -acv deploy cdn-release/
rsync -acv dist/ cdn-release/files/attribution-web/

echo $(date "+%Y%m%d%H%M%S")$RANDOM > cdn-release/random_for_zhiyun
