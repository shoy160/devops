<template>
  <a-layout>
    <a-layout-header>云智DevOps模板系统</a-layout-header>
    <a-layout-content>
      <a-form v-bind="formLayout" :form="form" @submit="handleSubmit">
        <a-form-item label="类型">
          <a-radio-group
            v-decorator="[
              'type',
              {
                initialValue: 'java',
                rules: [{ type: 'string', required: true }],
              },
            ]"
          >
            <a-radio-button value="java">
              Java
            </a-radio-button>
            <a-radio-button value="nodejs">
              Nodejs
            </a-radio-button>
          </a-radio-group>
        </a-form-item>
        <a-form-item label="Git地址">
          <a-input
            placeholder="请输入Git地址"
            v-decorator="[
              'git',
              {
                rules: [
                  { type: 'string', required: true, message: '请输入Git地址' },
                  {
                    validator: (rule, value) => validateGit(value),
                    message: 'Git地址不正确',
                  },
                ],
              },
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="镜像仓库">
          <a-select
            @change="handleGroupChange"
            v-decorator="[
              'imageGroup',
              {
                initialValue: '',
                rules: [
                  { type: 'string', required: true, message: '请选择镜像仓库' },
                ],
              },
            ]"
          >
            <a-select-option value="">-- 请选择 --</a-select-option>
            <a-select-option v-for="item in groups" :key="item.key">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="镜像名称">
          <a-input
            placeholder="请输入镜像名称"
            @change="handleImageChagne"
            v-decorator="[
              'image',
              {
                rules: [
                  { type: 'string', required: true, message: '请输入镜像名称' },
                ],
              },
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="云平台项目">
          <a-select
            v-decorator="[
              'project',
              {
                initialValue: '',
                rules: [
                  {
                    type: 'string',
                    required: true,
                    message: '请选择云平台项目',
                  },
                ],
              },
            ]"
          >
            <a-select-option value="">-- 请选择 --</a-select-option>
            <a-select-option v-for="item in projects" :key="item.key">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="云平台应用名称">
          <a-input
            placeholder="请输入云平台应用名称"
            v-decorator="[
              'app',
              {
                rules: [
                  {
                    type: 'string',
                    required: true,
                    message: '请输入云平台应用名称',
                  },
                ],
              },
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="云平台工作负载名称">
          <a-input
            placeholder="请输入云平台工作负载名称"
            v-decorator="[
              'workload',
              {
                rules: [
                  {
                    type: 'string',
                    required: true,
                    message: '请输入云平台工作负载名称',
                  },
                ],
              },
            ]"
          ></a-input>
        </a-form-item>
        <a-form-item label="云平台工作负载描述">
          <a-input
            placeholder="请输入云平台工作负载描述"
            v-decorator="['desc']"
          ></a-input>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" html-type="submit" icon="check-circle">
            确认提交
          </a-button>
        </a-form-item>
      </a-form>
    </a-layout-content>
    <a-layout-footer>Copyright © 2021</a-layout-footer>
  </a-layout>
</template>

<script>
import request from "@/utils/request"
export default {
  name: "Home",
  data() {
    return {
      formLayout: {
        labelCol: {
          xs: { span: 24 },
          sm: { span: 8 },
        },
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 16 },
        },
      },
      groups: [],
      projects: [],
      gitReg:
        "^https?:\\/\\/gitlab.yunzhicloud.com\\/(?<group>[^\\/]+)\\/(?<name>[^\\.]{2,})(\\.git)?$",
    }
  },
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "devops_form" })
  },
  mounted() {
    this.initGroups()
  },
  methods: {
    initGroups() {
      request.get("group").then((json) => {
        this.groups = json.data.groups
        this.projects = []
      })
    },
    validateGit(value) {
      var reg = new RegExp(this.gitReg, "ig")
      return reg.test(value)
    },
    handleGroupChange(e) {
      for (var i = 0; i < this.groups.length; i++) {
        const group = this.groups[i]
        if (group.key === e) {
          this.projects = group.projects
        }
      }
    },
    handleImageChagne(e) {
      var workload = this.form.getFieldValue("workload")
      var image = this.form.getFieldValue("image")
      if (workload && workload != image) {
        return
      }
      this.form.setFieldsValue({
        workload: e.target.value,
      })
    },
    handleSubmit(e) {
      e.preventDefault()
      this.form.validateFields((err, fieldsValue) => {
        if (err) {
          return
        }
        var reg = new RegExp(this.gitReg, "ig")
        var matches = reg.exec(fieldsValue["git"])
        if (!matches || !matches.groups) {
          return
        }
        const values = {
          ...fieldsValue,
          gitGroup: matches.groups.group,
          gitName: matches.groups.name,
        }
        delete values.git
        console.log("Received values of form: ", values)
        request.post("devops", values, { responseType: "blob" })
      })
    },
  },
}
</script>
<style lang="less">
.ant-layout-header {
  color: #fff;
  background-color: #24337d;
  text-align: center;
  font-size: 22px;
  line-height: 65px;
}
.ant-layout-content {
  max-width: 780px;
  min-width: 50%;
  margin: 50px auto;
}
</style>
