# Saga

- Consistency: our system is still in valid state after we have us when we've started the transaction and the end of transaction
- Atomicity
- Isolation: these transaction is not affect other if they are run concurrently
- Durability: commit something, it persist as long-lived and like we don't lose data

A Saga is a Collection of Sub-Transactions T1, T2, ... Tn
Each Sub-Transaction has a Compensating Transaction C1, C2 ... Cn
Cn Semantically Undoes Tn

Either: Saga Guarantee
- T1, T2, ... Tn or
- T1, T2, ... Tj, Cj, ... C2, C1

Sagas are a Failure Management Pattern

Saga Execution Coordinator (SEC). SEC lives in process with your database and so it sort of shares the same face so we're not in a distributed world. His job is to go and execute these sub transactions and then in the case of failure start applying compensating transactions.

Saga log that is just like your normal database log and so it's going to do things like you're going to commit the messages to it like begin saga, ends saga, abort saga and then all the beginning and commit message for the transactions and compensating transactions.

Unsuccessful Saga
- Backwards Recovery

A Distributed Saga is a Collection of Sub-Requests T1, T2, ... Tn
Each Sub-Request has a Compensating Request C1, C2, ... Cn

Successful Distributed Saga

Begin Saga

Start Book Hotel Request (T1)
End Book Hotel Request (T1)

Start Book Car Rental Request (T2)
End Book Car Rental Request (T2)

Start Book Flight Request (T3)
End Book Flight Request (T3)

End Saga

## Saga log
Durable & Distributed

We still need a log but now it's not co-located with a database because we don't have a single database we have to durable and distributed log that lives somewhere so this might be Kafka

## Saga Execution Coordinator (SEC)

This is once again i process and it doesn't live co-located 

Interprets & Writes  to Saga Log
Applies Saga Sub-Requests
Applies Saga Compensating Requests when Necessary


## Request Messaging Semantics

- Sub-Requests (Ti): At Most Once
- Compensating Requests (Ci): At Least Once

## Distributed Saga Guarantee
Either
- T1, T2, ... Tn or
- T1, T2, ... Tj Cj, ... C2, C1

## 2PC

- prepare phase: hỏi resource xem có sẵn sàng để thực hiện 1 transaction không
- nếu tất cả vote ok => commit phase
- execute phase