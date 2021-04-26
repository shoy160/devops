package com.yunzhicloud.network.dialtest.dao;

import com.baomidou.mybatisplus.core.metadata.IPage;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.yunzhicloud.network.dialtest.entity.dto.user.DepartmentDTO;
import com.yunzhicloud.network.dialtest.entity.po.sys.DepartmentPO;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.springframework.stereotype.Repository;

/**
 * @author codeGenerate
 */
@Repository
public interface DepartmentMapper extends BaseMapper<DepartmentPO> {
}
