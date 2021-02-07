package com.foodie;

import com.foodie.ordersvc.OrderServiceClient;
import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

@SpringBootApplication
@Slf4j
public class FoodieApiApplication {

  public static void main(String[] args) {
    SpringApplication.run(FoodieApiApplication.class, args);
  }

  @Bean
  CommandLineRunner run(OrderServiceClient client) {
    return args -> {
      //client.createOrder("barshatrai", "1011D", 101.0f);
      //var order = client.retrieveOrder(1L);
      //log.info(order.toString());
      var orderDtos = client.retrieveOrdersByUserID("barshatrai");
      for (var o : orderDtos) {
        log.info(o.toString());
      }
    };
  }
}
