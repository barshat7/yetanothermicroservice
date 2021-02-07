package com.foodie.ordersvc;

import com.foodie.dto.OrderDto;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import javax.annotation.PostConstruct;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import service.OrderServiceGrpc;
import service.OrderServiceOuterClass.Order;
import service.OrderServiceOuterClass.OrderID;
import service.OrderServiceOuterClass.OrderRequest;
import service.OrderServiceOuterClass.UserID;

@Component
@Slf4j
public class OrderServiceClient {

  @Value("${application.ordersvc.baseUrl}")
  private String orderServiceBaseUrl;

  @Value("${application.ordersvc.port}")
  private int orderServicePort;

  private  ManagedChannel channel;


  @PostConstruct
  public void init() {
    channel = ManagedChannelBuilder.forAddress(orderServiceBaseUrl, orderServicePort).usePlaintext().build();
  }

  public void createOrder(String userId, String itemId, float price) {
    var stub = OrderServiceGrpc.newBlockingStub(channel);
    var orderRequest =
        OrderRequest.newBuilder()
            .setUserID(userId)
            .addAllItemIDList(List.of("1", "2"))
            .setOrderDate(new Date().toString())
            .setTotalPrice(price)
            .build();

    OrderID id = stub.createOrder(orderRequest);
    log.info("Order Created With Order ID {}", id);
  }

  public OrderDto retrieveOrder(Long orderID) {
    var stub = OrderServiceGrpc.newBlockingStub(channel);
    Order order =
        stub.retrieveOrder(OrderID.newBuilder().setValue(String.valueOf(orderID)).build());
    return OrderDto.builder()
        .id(orderID)
        .userID(order.getUserID())
        .orderDate(order.getOrderDate())
        .totalPrice(order.getTotalPrice())
        .build();
  }

  public List<OrderDto> retrieveOrdersByUserID(String userID) {
    var stub = OrderServiceGrpc.newBlockingStub(channel);
    var request = UserID.newBuilder().setValue(userID).build();
    var itr = stub.retrieveOrdersByUserID(request);
    List<OrderDto> dtoList = new ArrayList<>();
    while (itr.hasNext()) {
      var order = itr.next();
      dtoList.add(
          OrderDto.builder()
              .userID(order.getUserID())
              .orderDate(order.getOrderDate())
              .totalPrice(order.getTotalPrice())
              .build());
    }
    return dtoList;
  }
}
