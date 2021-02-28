package com.foodie;

import com.foodie.cartsvc.CartServiceClient;
import com.foodie.dto.CartDto;
import com.foodie.dto.CartItemDto;
import com.foodie.ordersvc.OrderServiceClient;
import java.util.Arrays;
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
  CommandLineRunner run(OrderServiceClient client, CartServiceClient cartServiceClient) {
    return args -> {
      // client.createOrder("barshatrai", "1011D", 101.0f);
      // var order = client.retrieveOrder(1L);
      // log.info(order.toString());
      CartItemDto dto1 = new CartItemDto(12L, "Bedsheets", 900.00, 10);
      CartItemDto dto2 = new CartItemDto(17L, "Smartphone", 1200.00, 10);
      CartItemDto dto3 = new CartItemDto(14L, "Groceries", 900.00, 10);
    //  CartDto cartDto = new CartDto();
//      cartDto.setUserID("reyo999");
//      cartDto.setCartItemDtoList( Arrays.asList(dto1, dto2, dto3));
      //cartServiceClient.addItemsToCart("reyo999", Arrays.asList(dto1, dto2, dto3));
      //cartServiceClient.retrieveCartByUserID("reyo99");
    };
  }
}
