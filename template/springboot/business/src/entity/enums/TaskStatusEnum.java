package com.yunzhicloud.network.dialtest.enums;

import com.yunzhicloud.core.enums.ValueNameEnum;
import lombok.Getter;

/**
 * 任务告警状态
 * 0故障  1恢复
 *
 * @author pengwei
 * @date 2021/3/11
 */
@Getter
public enum TaskStatusEnum implements ValueNameEnum<Integer> {

    FAULT(0, "故障"),
    RECOVER(1, "恢复");
    private final int code;
    private final String desc;

    TaskStatusEnum(int code, String desc) {
        this.code = code;
        this.desc = desc;

    }

    public static String getDescByCode(int code) {
        for (TaskStatusEnum enu : TaskStatusEnum.values()) {
            if (enu.getCode() == code) {
                return enu.getDesc();
            }
        }
        return "";
    }
    @Override
    public String getName() {
        return this.desc;
    }
    @Override
    public Integer getValue() {
        return this.code;
    }
}
