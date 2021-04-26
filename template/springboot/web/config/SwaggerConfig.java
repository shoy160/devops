package com.yunzhicloud.network.dialtest.web.config;

import com.github.xiaoymin.knife4j.spring.annotations.EnableKnife4j;
import com.yunzhicloud.network.dialtest.DialTestConstants;
import com.yunzhicloud.web.swagger.BaseSwaggerConfig;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import springfox.bean.validators.configuration.BeanValidatorPluginsConfiguration;
import springfox.documentation.builders.ParameterBuilder;
import springfox.documentation.schema.ModelRef;
import springfox.documentation.service.Parameter;
import springfox.documentation.spring.web.plugins.Docket;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

import java.util.ArrayList;
import java.util.List;

/**
 * Todo
 *
 * @author shay
 * @date 2020/7/21
 */
@Configuration
@EnableSwagger2
@EnableKnife4j
@Import(BeanValidatorPluginsConfiguration.class)
public class SwaggerConfig extends BaseSwaggerConfig {
    @Bean
    public Docket propertyAppletDocket() {
        ParameterBuilder ticketPar = new ParameterBuilder();
        List<Parameter> pars = new ArrayList<>();
        ticketPar.name("Jwt-Token").description("Jwt-Token认证")
                .modelRef(new ModelRef("string")).parameterType("header")
                //header中的ticket参数非必填，传空也可以
                .required(false).build();
        //根据每个方法名也知道当前方法在设置什么参数
        pars.add(ticketPar.build());

        return getDocket("拨测系统", DialTestConstants.VERSION, DialTestConstants.REST_PACKAGE)
                .globalOperationParameters(pars);
    }
}
