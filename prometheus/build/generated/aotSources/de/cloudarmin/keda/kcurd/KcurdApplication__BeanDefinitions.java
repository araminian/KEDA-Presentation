package de.cloudarmin.keda.kcurd;

import org.springframework.aot.generate.Generated;
import org.springframework.beans.factory.config.BeanDefinition;
import org.springframework.beans.factory.support.RootBeanDefinition;
import org.springframework.context.annotation.ConfigurationClassUtils;

/**
 * Bean definitions for {@link KcurdApplication}.
 */
@Generated
public class KcurdApplication__BeanDefinitions {
  /**
   * Get the bean definition for 'kcurdApplication'.
   */
  public static BeanDefinition getKcurdApplicationBeanDefinition() {
    RootBeanDefinition beanDefinition = new RootBeanDefinition(KcurdApplication.class);
    beanDefinition.setTargetType(KcurdApplication.class);
    ConfigurationClassUtils.initializeConfigurationClass(KcurdApplication.class);
    beanDefinition.setInstanceSupplier(KcurdApplication$$SpringCGLIB$$0::new);
    return beanDefinition;
  }
}
