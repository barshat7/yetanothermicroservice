package com.foodie.cartsvc;

import cartgrpc.CartServiceOuterClass.AddCartItemRequest;
import cartgrpc.CartServiceOuterClass.CartItemRequest;
import cartgrpc.CartServiceOuterClass.CreateCartRequest;
import cartgrpc.CartServiceOuterClass.UserID;
import com.foodie.dto.CartDto;
import com.foodie.dto.CartItemDto;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.util.ArrayList;
import java.util.List;
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
        .setQuantity(item.getQuantity())
        .build())
        .collect(Collectors.toList());
    var createCartRequest = CreateCartRequest.newBuilder()
        .setUserID(cartDto.getUserID())
        .addAllCartItems(allCartItems)
        .build();
    cartgrpc.CartServiceOuterClass.CartID cartID = stub.createCart(createCartRequest);
    log.info("Cart Created For User {} With Cart ID {}", cartDto.getUserID(), cartID.getValue());
  }

  public void retrieveCartByUserID(String userID) {
    var stub = cartgrpc.CartServiceGrpc.newBlockingStub(channel);
    var userIDRequest = UserID.newBuilder().setValue(userID).build();
    var cartResponse = stub.retrieveCartByUserID(userIDRequest);
    log.info("Cart Created Date {}", cartResponse.getCreatedDate());
    log.info("Items in cart");
    cartResponse.getCartItemsList().forEach(item -> log.info(item.getItemName()));
  }

  public void addItemsToCart(String userID, List<CartItemDto> cartItemDtoList) {
    var stub = cartgrpc.CartServiceGrpc.newBlockingStub(channel);
    var userIDRequest = UserID.newBuilder().setValue(userID).build();
    var addCartItemRequest = AddCartItemRequest.newBuilder()
        .addAllCartItems(cartItemDtoList.stream()
        .map(cartItemDto -> CartItemRequest.newBuilder()
        .setItemID(cartItemDto.getItemID())
        .setItemName(cartItemDto.getItemName())
        .setItemPrice(cartItemDto.getItemPrice())
        .setQuantity(cartItemDto.getQuantity())
        .build())
            .collect(Collectors.toList()))
        .setUserID(userID)
        .build();
   var cartID = stub.addItemToCart(addCartItemRequest);
   log.info("Items Added To Cart ID {}", cartID.getValue());
  }
}
