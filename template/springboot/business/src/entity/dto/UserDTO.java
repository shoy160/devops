package com.yunzhicloud.network.dialtest.entity.dto.user;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.Data;
import lombok.experimental.Accessors;

import java.io.Serializable;
import java.util.Date;

/**
 * @author codeGenerate
 */
@Data
@Accessors(chain = true)
@JsonInclude(JsonInclude.Include.NON_NULL)
public class UserDTO implements Serializable{

    private static final long serialVersionUID = 1L;

    /** 主键id*/
    private Integer id;

    /** 成员编号*/
    private String num;

    /** 用户名*/
    private String name;

    /** 帐号状态（0正常 1停用）*/
    private Integer status;

    /** 用户昵称*/
    private String nickName;

    /** 盐值*/
    private String salt;

    /** 手机号*/
    private String phone;

    /** 密码*/
    private String password;

    /** 头像地址*/
    private String avatar;

    /** 用户性别(0男 1女 2未知)*/
    private Integer sex;

    /** 邮箱地址*/
    private String email;

    /** 成员身份(0普通 1部门领导 2高层领导*/
    private Integer identity;

    /** 部门id*/
    private Integer departmentId;

    /** 入职日期*/
    private Date entryDate;

    /** 最后登陆IP地址*/
    private String loginIp;

    /** 最后登陆时间*/
    private Date loginDate;

    /** 离职时间*/
    private Date resignDate;

    /** 是否离职 1是 0 否*/
    private Integer isResign;

    /** 备注*/
    private String remark;

    /** 是否删除 1是 0否*/
    private Integer deleted;

    /** 创建时间*/
    private Date createDate;

    /** 更新时间*/
    private Date updateDate;

    /** 保留字段1*/
    private String reserved1;

    /** 保留字段2*/
    private String reserved2;

}
