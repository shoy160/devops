# coding=utf-8
import sys
import getopt
import json
import random
import shutil
import os


def usage():
    help = 'main.py -t[--type=]\t\tDevOps类型(java|nodejs)\n'
    help = help + '\t-a[--image-group=]\t镜像仓库组\n'
    help = help + '\t-b[--image-name=]\t镜像名称\n'
    help = help + '\t-c[--git-group=]\tgitlab分组\n'
    help = help + '\t-d[--git-name=]\t\tgitlab项目名\n'
    help = help + '\t-e[--namespace=]\t云平台项目名称\n'
    help = help + '\t-f[--app-name=]\t\t云平台应用名称\n'
    help = help + '\t-g[--deploy-name=]\t云平台工作负载名称\n'
    help = help + '\t-i[--deploy-desc=]\t云平台工作负载描述\n'
    print(help)


def get_params(args, required=True):
    cnf = {}
    try:
        opts, args = getopt.getopt(
            args, "ha:b:c:d:e:f:g:i:t:",
            ["help", "image-group=", "image-name=", "git-group=", "git-name=", "namespace=", "app-name=", "deploy-name=", "deploy-desc=", "type="])
    except getopt.GetoptError:
        usage()
        sys.exit(2)
    for opt, arg in opts:
        # print(opt, arg)
        if opt == '-h' or opt == '--help':
            usage()
            sys.exit()
        if opt == '-a' or opt == '--image-group':
            cnf['image_group'] = arg
        if opt == '-b' or opt == '--image-name':
            cnf['image_name'] = arg
        if opt == '-c' or opt == '--git-group':
            cnf['git_group'] = arg
        if opt == '-d' or opt == '--git-name':
            cnf['git_name'] = arg
        if opt == '-e' or opt == '--namespace':
            cnf['ns_name'] = arg
        if opt == '-f' or opt == '--app-name':
            cnf['app_name'] = arg
        if opt == '-g' or opt == '--deploy-name':
            cnf['deploy_name'] = arg
        if opt == '-i' or opt == '--deploy-desc':
            cnf['deploy_desc'] = arg
        if opt == '-t' or opt == '--type':
            cnf['type'] = arg
    if not required:
        return cnf
    if 'type' not in cnf or not cnf['type']:
        cnf['type'] = input('DevOps类型(java|nodejs)：')
    if 'image_group' not in cnf or not cnf['image_group']:
        cnf['image_group'] = input('镜像仓库组名：')
    if 'image_name' not in cnf or not cnf['image_name']:
        cnf['image_name'] = input('镜像名称：')
    if 'git_group' not in cnf or not cnf['git_group']:
        cnf['git_group'] = input('gitlab分组：')
    if 'git_name' not in cnf or not cnf['git_name']:
        cnf['git_name'] = input('gitlab项目名：')
    if 'ns_name' not in cnf or not cnf['ns_name']:
        cnf['ns_name'] = input('云平台项目名称：')
    if 'app_name' not in cnf or not cnf['app_name']:
        cnf['app_name'] = input('云平台应用名称：')
    if 'deploy_name' not in cnf or not cnf['deploy_name']:
        cnf['deploy_name'] = input('云平台工作负载名称：')
    if 'deploy_desc' not in cnf or not cnf['deploy_desc']:
        cnf['deploy_desc'] = input('云平台工作负载描述：')
    return cnf


def replace_content(content, cnf):
    for name in cnf:
        if name == 'type':
            continue
        variable = '${%s}' % name.replace('_', '-')
        content = content.replace(variable, cnf[name])
    prod_ns = cnf['ns_name']
    if 'test' in prod_ns:
        prod_ns = prod_ns.replace('test', 'prod')
    else:
        prod_ns = '%s-prod' % prod_ns
    content = content.replace('${ns-name-prod}', prod_ns)
    return content


def replace_files(path, cnf):
    path_list = os.listdir(path)
    for file in path_list:
        abs_path = os.path.join(path, file)
        if os.path.isfile(abs_path):
            with open(abs_path, 'r', encoding='utf-8') as f:
                file_content = f.read()
            with open(abs_path, 'w', encoding='utf-8') as f:
                f.write(replace_content(file_content, cnf))
        else:
            replace_files(abs_path, cnf)


def make_devops(cnf):
    print(cnf)
    confirm = input('确认DevOps信息(y|n):')
    if confirm == 'y':
        # dir = cnf['type']
        # copy
        devops_name = '%s_%d' % (cnf['type'], random.randint(1000, 10000))
        source_dir = os.path.abspath(cnf['type'])
        devops_dir = os.path.abspath(devops_name)
        shutil.copytree(source_dir, devops_dir)
        print('copy files to %s' % devops_name)
        # 替换字符
        replace_files(devops_dir, cnf)


if __name__ == "__main__":
    cnf = get_params(sys.argv[1:])
    make_devops(cnf)
