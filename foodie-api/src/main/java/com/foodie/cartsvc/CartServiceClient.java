package com.foodie.cartsvc;

import cartgrpc.CartServiceOuterClass.CartItemRequest;
import cartgrpc.CartServiceOuterClass.CreateCartRequest;
import com.foodie.dto.CartDto;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.util.ArrayList;
import java.util.stream.Collectors;
import javax.annotation.PostConstruct;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
@Slf4j
public class CartServiceClient {

  @Value("${application.cartsvc.baseUrl}")
  private String cartServiceBaseUrl;

  @Value("${application.cartsvc.port}")
  private int cartServicePort;

  private ManagedChannel channel;


  @PostConstruct
  public void init() {
    channel = ManagedChannelBuilder.forAddress(cartServiceBaseUrl, cartServicePort).usePlaintext().build();
  }

  public void createCart(CartDto cartDto) {
    var stub = cartgrpc.CartServiceGrpc.newBlockingStub(channel);
    var allCartItems = cartDto.getCartItemDtoList().stream().map(item -> CartItemRequest.newBuilder()
        .setItemID(item.getItemID())
        .setItemName(item.getItemName())
        .setItemPrice(item.getItemPrice())
        .build())
        .collect(Collectors.toList());
    var createCartRequest = CreateCartRequest.newBuilder()
        .setUserID(cartDto.getUserID())
        .addAllCartItems(allCartItems)
        .build();
    cartgrpc.CartServiceOuterClass.CartID cartID = stub.createCart(createCartRequest);
    log.info("Cart Created For User {} With Cart ID {}", cartDto.getUserID(), cartID.getValue());
  }
}
