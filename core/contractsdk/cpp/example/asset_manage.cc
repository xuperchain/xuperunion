#include "xchain/xchain.h"

const std::string BALANCEPRE = "balanceOf_";
const std::string ALLOWANCEPRE = "allowanceOf_";
const std::string MASTERPRE = "owner";

// 资产管理合约的基类
// 资产管理合约需要实现基类中指定的方法
// 参数由xchain::Contract中的context提供
class AssetBasic {
public:
    /*
     * func: 初始化资产管理账户以及总发行量
     * @param: initiator:交易发起者,也是初始化资产的owner
     * @param: totalSupply:发行总量,初始化时,资产全部归initiator
     */
    virtual void initialize() = 0;
    /*
     * func: 增发资产
      * @param: initiator:交易发起者,只有交易发起者等于资产owner时，才能增发
      * @param: amount:增发容量
      */
    virtual void mint() = 0;
    /*
     * func: 获取资产总供应量
     */
    virtual void totalSupply() = 0;
    /*
     * func: 获取caller的资产余额
     * @param: caller: 合约调用者
     */
    virtual void balance() = 0;
    /*
     * func: 查询to用户能消费from用户的资产数量
     * @param: from: 被消费方
     * @param: to: 消费方
     */
    virtual void allowance() = 0;
    /*
     * func: from账户给to账户转token数量的资产
     * @param: from:转账方
     * @param: to:收账方
     * @param: token:转账资产数量
     */
    virtual void transfer() = 0;
    /*
     * func: 从授权账户from转移数量为token的资产给to账户
      * @param: from:被转账账户
      * @param: caller:合约调用者
      * @param: to:收账账户
      * @param: token:转移的资产数量
      */
    virtual void transferFrom() = 0;
    /*
      * func: 允许to账户从from账户转移token数量的资产
     * @param: from:
     * @param: to:
     * @param: token
     */
    virtual void approve() = 0;
};

struct ERC20 : public AssetBasic, public xchain::Contract {
public:
    void initialize() {
        xchain::Context* ctx = this->context();
        const std::string& caller = ctx->initiator();
        if (caller.empty()) {
            ctx->error("missing caller");
            return;
        }
        const std::string& totalSupply = ctx->arg("totalSupply");
        if (totalSupply.empty()) {
            ctx->error("missing totalSupply");
            return;
        }

        std::string key = BALANCEPRE + caller;
        ctx->put_object("totalSupply", totalSupply);
        ctx->put_object(key, totalSupply);

        std::string master = MASTERPRE;
        ctx->put_object(master, caller);
        ctx->ok("initialize success");
    }
    void mint() {
        xchain::Context* ctx = this->context();
        const std::string& caller = ctx->initiator();
        if (caller.empty()) {
            ctx->error("missing caller");
            return;
        }

        std::string master;
        if (!ctx->get_object(MASTERPRE, &master)) {
            ctx->error("missing master");
            return;
        }
        if (master != caller) {
            ctx->error("only the person who created the contract can mint");
            return;
        }

        const std::string& increaseSupply = ctx->arg("amount");
        if (increaseSupply.empty()) {
            ctx->error("missing amount");
            return;
        }

        std::string value;
        if (!ctx->get_object("totalSupply", &value)) {
            ctx->error("get totalSupply error");
            return;
        }
        int increaseSupplyint = atoi(increaseSupply.c_str());
        int valueint = atoi(value.c_str());
        int totalSupplyint = increaseSupplyint + valueint;
        char buf[32];
        snprintf(buf, 32, "%d", totalSupplyint);
        ctx->put_object("totalSupply", buf); 

        std::string key = BALANCEPRE + caller;
        if (!ctx->get_object(key, &value)) {
            ctx->error("get caller balance error");
            return;
        }
        valueint = atoi(value.c_str());
        int callerint = increaseSupplyint + valueint;
        snprintf(buf, 32, "%d", callerint);
        ctx->put_object(key, buf); 
    
        ctx->ok(buf);
    }
    void totalSupply() {
        xchain::Context* ctx = this->context();
        std::string value;
        if (ctx->get_object("totalSupply", &value)) {
            ctx->ok(value);
        } else {
            ctx->error("key not found");
        }
    }
    void balance() {
        xchain::Context* ctx = this->context();
        const std::string& caller = ctx->arg("caller");
        if (caller.empty()) {
            ctx->error("missing caller");
            return;
        }

        std::string key = BALANCEPRE + caller;
        std::string value;
        if (ctx->get_object(key, &value)) {
            ctx->ok(value);
        } else {
            ctx->error("key not found");
        }
    }
    void allowance() {
        xchain::Context* ctx = this->context();
        const std::string& from = ctx->arg("from");
        if (from.empty()) {
            ctx->error("missing from");
            return;
        }

        const std::string& to = ctx->arg("to");
        if (to.empty()) {
            ctx->error("missing to");
            return;
        }

        std::string key = ALLOWANCEPRE + from + "_" + to;
        std::string value;
        if (ctx->get_object(key, &value)) {
            ctx->ok(value);
        } else {
            ctx->error("key not found");
        }
    }
    void transfer() {
        xchain::Context* ctx = this->context();
        const std::string& from = ctx->arg("from");
        if (from.empty()) {
            ctx->error("missing from");
            return;
        }
   
        const std::string& to = ctx->arg("to");
        if (to.empty()) {
            ctx->error("missing to");
            return;
        }

        const std::string& token_str = ctx->arg("token");
        if (token_str.empty()) {
            ctx->error("missing token");
            return;
        }
        int token = atoi(token_str.c_str());

        std::string from_key = BALANCEPRE + from;
        std::string value;
        int from_balance = 0;
        if (ctx->get_object(from_key, &value)) {
            from_balance = atoi(value.c_str()); 
            if (from_balance < token) {
                ctx->error("The balance of from not enough");
                return;
            }  
        } else {
            ctx->error("key not found");
            return;
        }

        std::string to_key = BALANCEPRE + to;
        int to_balance = 0;
        if (ctx->get_object(to_key, &value)) {
            to_balance = atoi(value.c_str());
        }
   
        from_balance = from_balance - token;
        to_balance = to_balance + token;
   
        char buf[32]; 
        snprintf(buf, 32, "%d", from_balance);
        ctx->put_object(from_key, buf);
        snprintf(buf, 32, "%d", to_balance);
        ctx->put_object(to_key, buf);

        ctx->ok("transfer success");
    }
    void transferFrom() {
        xchain::Context* ctx = this->context();
        const std::string& from = ctx->arg("from");
        if (from.empty()) {
            ctx->error("missing from");
            return;
        }
  
        const std::string& caller = ctx->arg("caller");
        if (caller.empty()) {
            ctx->error("missing caller");
            return;
        }

        const std::string& to = ctx->arg("to");
        if (to.empty()) {
            ctx->error("missing to");
            return;
        }

        const std::string& token_str = ctx->arg("token");
        if (token_str.empty()) {
            ctx->error("missing token");
            return;
        }
        int token = atoi(token_str.c_str());

        std::string allowance_key = ALLOWANCEPRE + from + "_" + caller;
        std::string value;
        int allowance_balance = 0;
        if (ctx->get_object(allowance_key, &value)) {
            allowance_balance = atoi(value.c_str()); 
            if (allowance_balance < token) {
                ctx->error("The allowance of from_to not enough");
                return;
            }  
        } else {
            ctx->error("You need to add allowance from_to");
            return;
        }

        std::string from_key = BALANCEPRE + from;
        int from_balance = 0;
        if (ctx->get_object(from_key, &value)) {
            from_balance = atoi(value.c_str()); 
            if (from_balance < token) {
                ctx->error("The balance of from not enough");
                return;
            }  
        } else {
            ctx->error("From no balance");
            return;
        }

        std::string to_key = BALANCEPRE + to;
        int to_balance = 0;
        if (ctx->get_object(to_key, &value)) {
            to_balance = atoi(value.c_str());
        }
   
        from_balance = from_balance - token;
        to_balance = to_balance + token;
        allowance_balance = allowance_balance - token;

        char buf[32]; 
        snprintf(buf, 32, "%d", from_balance);
        ctx->put_object(from_key, buf);
        snprintf(buf, 32, "%d", to_balance);
        ctx->put_object(to_key, buf);
        snprintf(buf, 32, "%d", allowance_balance);
        ctx->put_object(allowance_key, buf);

        ctx->ok("transferFrom success");
    }
    void approve() {
        xchain::Context* ctx = this->context();
        const std::string& from = ctx->arg("from");
        if (from.empty()) {
            ctx->error("missing from");
            return;
        }
   
        const std::string& to = ctx->arg("to");
        if (to.empty()) {
            ctx->error("missing to");
            return;
        }

        const std::string& token_str = ctx->arg("token");
        if (token_str.empty()) {
            ctx->error("missing token");
            return;
        }
        int token = atoi(token_str.c_str());

        std::string from_key = BALANCEPRE + from;
        std::string value;
        if (ctx->get_object(from_key, &value)) {
            int from_balance = atoi(value.c_str()); 
            if (from_balance < token) {
                ctx->error("The balance of from not enough");
                return;
            }  
        } else {
            ctx->error("From no balance");
            return;
        }

        std::string allowance_key = ALLOWANCEPRE + from + "_" + to;
        int allowance_balance = 0;
        if (ctx->get_object(allowance_key, &value)) {
            allowance_balance = atoi(value.c_str());
        }

        allowance_balance = allowance_balance + token;
   
        char buf[32]; 
        snprintf(buf, 32, "%d", allowance_balance);
        ctx->put_object(allowance_key, buf);

        ctx->ok("approve success");
    }
};


DEFINE_METHOD(ERC20, initialize) {
    self.initialize();
}

DEFINE_METHOD(ERC20, mint) {
    self.mint();
}

DEFINE_METHOD(ERC20, totalSupply) {
    self.totalSupply();
}

DEFINE_METHOD(ERC20, balance) {
    self.balance();
}

DEFINE_METHOD(ERC20, allowance) {
    self.allowance();
}

DEFINE_METHOD(ERC20, transfer) {
    self.transfer();
}

DEFINE_METHOD(ERC20, transferFrom) {
    self.transferFrom();
}

DEFINE_METHOD(ERC20, approve) {
    self.approve();
}
