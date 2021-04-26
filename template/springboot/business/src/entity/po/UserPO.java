package com.yunzhicloud.network.dialtest.entity.po.user;

import com.baomidou.mybatisplus.annotation.*;
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
@TableName("sys_user")
@JsonInclude(JsonInclude.Include.NON_NULL)
public class UserPO implements Serializable{

    private static final long serialVersionUID = 1L;

    /** 主键id*/
    @TableId(type = IdType.AUTO,value = "fd_id")
    private Integer id;

    /** 成员编号*/
    @TableField("fd_num")
    private String num;

    /** 用户名*/
    @TableField("fd_name")
    private String name;

    /** 帐号状态（0正常 1停用）*/
    @TableField("fd_status")
    private Integer status;

    /** 用户昵称*/
    @TableField("fd_nickName")
    private String nickName;

    /** 盐值*/
    @TableField("fd_salt")
    private String salt;

    /** 手机号*/
    @TableField("fd_phone")
    private String phone;

    /** 密码*/
    @TableField("fd_password")
    private String password;

    /** 头像地址*/
    @TableField("fd_avatar")
    private String avatar;

    /** 用户性别(0男 1女 2未知)*/
    @TableField("fd_sex")
    private Integer sex;

    /** 邮箱地址*/
    @TableField("fd_email")
    private String email;

    /** 成员身份(0普通 1部门领导 2高层领导*/
    @TableField("fd_identity")
    private Integer identity;

    /** 部门id*/
    @TableField("fd_departmentId")
    private Integer departmentId;

    /** 入职日期*/
    @TableField("fd_entryDate")
    private Date entryDate;

    /** 最后登陆IP地址*/
    @TableField("fd_loginIp")
    private String loginIp;

    /** 最后登陆时间*/
    @TableField("fd_loginDate")
    private Date loginDate;

    /** 离职时间*/
    @TableField("fd_resignDate")
    private Date resignDate;

    /** 是否离职 1是 0 否*/
    @TableField("fd_isResign")
    private Integer isResign;

    /** 备注*/
    @TableField("fd_remark")
    private String remark;

    /** 是否删除 1是 0否*/
    @TableLogic
    @TableField("fd_deleted")
    private Integer deleted;

    /** 创建时间*/
    @TableField("fd_createDate")
    private Date createDate;

    /** 更新时间*/
    @TableField("fd_updateDate")
    private Date updateDate;

    /** 保留字段1*/
    @TableField("fd_reserved1")
    private String reserved1;

    /** 保留字段2*/
    @TableField("fd_reserved2")
    private String reserved2;

}
