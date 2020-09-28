# cd 到 cdn-release 目录，该目录在 cdn-release.sh 执行时生成
CURRENT_DIR=$(cd $(dirname $0); pwd)
cd ${CURRENT_DIR}
cd ../

if [ -f "/usr/local/services/cdn_rsync-1.0/bin/cdn_rsync" ]
then
    echo '同步 cdn'
    /usr/local/services/cdn_rsync-1.0/bin/cdn_rsync -acvz --port=8730 --bwlimit=1200 --timeout=60 files/ root@cdn_l5::ams_web/
else
    echo 'rsync -acvz ./files/ rsync://user_00@10.241.87.150:8730/ams_web'
    # 将 deploy 文件夹及多个静态文件夹迁移到 cdn-release 文件夹下
    rsync -acvz files/ rsync://user_00@10.241.87.150:8730/ams_web
    rsync -acvz files/ rsync://user_00@10.101.103.106:8730/ams_web
fi