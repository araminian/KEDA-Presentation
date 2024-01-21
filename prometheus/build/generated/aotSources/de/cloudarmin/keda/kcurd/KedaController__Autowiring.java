package de.cloudarmin.keda.kcurd;

import io.micrometer.core.instrument.MeterRegistry;
import org.springframework.aot.generate.Generated;
import org.springframework.beans.factory.aot.AutowiredMethodArgumentsResolver;
import org.springframework.beans.factory.support.RegisteredBean;

/**
 * Autowiring for {@link KedaController}.
 */
@Generated
public class KedaController__Autowiring {
  /**
   * Apply the autowiring.
   */
  public static KedaController apply(RegisteredBean registeredBean, KedaController instance) {
    AutowiredMethodArgumentsResolver.forRequiredMethod("setMetrics", MeterRegistry.class).resolve(registeredBean, args -> instance.setMetrics(args.get(0)));
    return instance;
  }
}
