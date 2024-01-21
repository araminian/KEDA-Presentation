package de.cloudarmin.keda.kcurd;

import org.springframework.aot.generate.Generated;
import org.springframework.beans.factory.config.BeanDefinition;
import org.springframework.beans.factory.support.InstanceSupplier;
import org.springframework.beans.factory.support.RootBeanDefinition;

/**
 * Bean definitions for {@link KedaController}.
 */
@Generated
public class KedaController__BeanDefinitions {
  /**
   * Get the bean definition for 'kedaController'.
   */
  public static BeanDefinition getKedaControllerBeanDefinition() {
    RootBeanDefinition beanDefinition = new RootBeanDefinition(KedaController.class);
    InstanceSupplier<KedaController> instanceSupplier = InstanceSupplier.using(KedaController::new);
    instanceSupplier = instanceSupplier.andThen(KedaController__Autowiring::apply);
    beanDefinition.setInstanceSupplier(instanceSupplier);
    return beanDefinition;
  }
}
