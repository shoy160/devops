import { message } from "ant-design-vue"
import Axios from "axios"

Axios.defaults.withCredentials = true

const messageKey = "msg-global"
const baseApi = process.env.VUE_APP_API
const request = Axios.create({
  baseURL: baseApi,
  timeout: 5000,
})
request.url = (api) => {
  return `${baseApi}/${api}`
}
request.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    console.error(error)
    Promise.reject(error)
  }
)

// respone interceptor
request.interceptors.response.use(
  (response) => {
    try {
      console.log(response)
      if (response.config.responseType === "blob") {
        const blob = new Blob([response.data], {
          type: response.headers["content-type"],
        })
        let contentDisposition = response.headers["content-disposition"]
        let patt = new RegExp("filename=[\"']?([^;]+\\.[^\\.;\"']+)[\"']?;*")
        let result = patt.exec(contentDisposition)
        let filename = decodeURI(result[1]).trim()
        const downloadElement = document.createElement("a")
        if ("download" in downloadElement) {
          const href = window.URL.createObjectURL(blob)
          downloadElement.style.display = "none"
          downloadElement.href = href
          downloadElement.download = filename
          document.body.appendChild(downloadElement)
          downloadElement.click()
          document.body.removeChild(downloadElement)
          window.URL.revokeObjectURL(href)
        } else {
          navigator.msSaveBlob(blob, filename)
        }
        message.success("文件下载成功！")
        return {}
      }
      var res = response.data
      if (!res.success) {
        //错误码处理
        message.error({
          content: res.message,
          duration: 2,
          key: messageKey,
        })
        return Promise.reject({
          status: res.code || 500,
          message: res.message,
        })
      }
      return res
    } catch (e) {
      return Promise.reject(e)
    }
  },
  (error) => {
    console.error(error)
    message.error({
      content: "网络异常，请稍后重试",
      duration: 2,
      key: messageKey,
    })
    return Promise.reject(error)
  }
)

export const getUrl = (api) => {
  return `${baseApi}/${api}`
}

export default request
